package Array

import "math"

//面试题 16.03. 交点
func intersection(start1 []int, end1 []int, start2 []int, end2 []int) []float64 {
	Ax, Ay, Bx, By, Cx, Cy, Dx, Dy := start1[0], start1[1], end1[0], end1[1], start2[0], start2[1], end2[0],end2[1]

	if IsInSameLineOrTShape(Ax, Ay, Bx, By, Cx, Cy, Dx, Dy) {
		aInCD := BothSide(Ax, Ay, Cx, Cy, Dx, Dy)
		bInCD := BothSide(Bx, By, Cx, Cy, Dx, Dy)
		cInAB := BothSide(Cx, Cy, Ax, Ay, Bx, By)
		dInAB := BothSide(Dx, Dy, Ax, Ay, Bx, By)
		var retpoint []float64
		if aInCD { Updata(&retpoint,[]float64{float64(Ax),float64(Ay)}) }
		if bInCD { Updata(&retpoint,[]float64{float64(Bx),float64(By)}) }
		if cInAB { Updata(&retpoint,[]float64{float64(Cx),float64(Cy)}) }
		if dInAB { Updata(&retpoint,[]float64{float64(Dx),float64(Dy)}) }
		return retpoint
	}

	if !HasIntersection(Ax, Ay, Bx, By, Cx, Cy, Dx, Dy) {
		return []float64{}
	}
	//计算交点
	sABC := Areas(Ax, Ay, Bx, By, Cx, Cy)
	sABD := Areas(Ax, Ay, Bx, By, Dx, Dy)
	k := sABC/sABD
	x := (float64(Cx) + k * float64(Dx)) / (1+k)
	y := (float64(Cy) + k * float64(Dy)) / (1+k)
	return []float64{x,y}
}

//向量叉乘
//几何意义：二维向量叉乘的绝对值几何上表示两个向量形成的平行四边形的面积
//那么向量组成的三角形面积自然就是 1/2 * S
func cross(x1,y1,x2,y2 int) int {
	return x1*y2 - x2*y1
}
//三角形的面积
func Areas(x1,y1,x2,y2,x3,y3 int) float64 {
	X1 := x3 - x1
	Y1 := y3 - y1
	X2 := x2 - x1
	Y2 := y2 - y1
	return math.Abs(0.5*float64(cross(X1,Y1,X2,Y2)))
}
//判断共线或者呈现T型
func IsInSameLineOrTShape(x1,y1, x2,y2, x3,y3, x4,y4 int) bool {
	ABx := x2 - x1
	ABy := y2 - y1
	cx := x3 - x1
	cy := y3 - y1
	dx := x4 - x1
	dy := y4 - y1

	if cross(ABx,ABy,cx,cy) == 0 || cross(ABx,ABy,dx,dy) == 0 {
		return true
	}
	return false
}
//判断共线之后,判断点是否在点线段上
func BothSide(x1,y1,x2,y2,x3,y3 int) bool  {
	return (x2-x1)*(x3-x1) <= 0 && (y2-y1)*(y3-y1) <= 0
}
//判断2个线段是否有交点
func HasIntersection(x1,y1, x2,y2, x3,y3, x4,y4 int) bool {
	ABx := x2 - x1
	ABy := y2 - y1
	cx := x3 - x1
	cy := y3 - y1
	dx := x4 - x1
	dy := y4 - y1
	if cross(ABx,ABy,cx,cy) * cross(ABx,ABy,dx,dy) > 0 {
		return false
	}

	CDx := x4 - x3
	CDy := y4 - y3
	ax := x1 - x3
	ay := y1 - y3
	bx := x2 - x3
	by := y2 - y3
	if cross(CDx,CDy,ax,ay) * cross(CDx,CDy,bx,by) > 0 {
		return false
	}

	return true
}

func Updata(ret *[]float64, newpoint []float64)  {
	if len(newpoint) == 0 {
		return
	}

	if len(*ret) == 0 {
		for i := 0 ; i < len(newpoint) ; i++ {
			*ret = append(*ret,newpoint[i])
		}
		return
	}

	if (*ret)[0] > newpoint[0] {
		for i := 0 ; i < len(newpoint) ; i++ {
			*ret = append(*ret,newpoint[i])
		}
		return
	}

	if (*ret)[0] == newpoint[0] && (*ret)[1] > newpoint[1]  {
		for i := 0 ; i < len(newpoint) ; i++ {
			*ret = append(*ret,newpoint[i])
		}
		return
	}
}
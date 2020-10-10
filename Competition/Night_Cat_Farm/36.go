package Night_Cat_Farm

//1603. 设计停车系统
type ParkingSystem struct {
	big int
	mid int
	small int
}


func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		big : big,
		mid : medium,
		small : small,
	}
}


func (this *ParkingSystem) AddCar(carType int) bool {
	if carType == 1 {
		if this.big > 0 {
			this.big--
			return true
		}else {
			return false
		}
	}else if carType == 2 {
		if this.mid > 0 {
			this.mid--
			return true
		}else {
			return false
		}
	}else if carType == 3 {
		if this.small > 0 {
			this.small--
			return true
		}else {
			return false
		}
	}
	return false
}


/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */

//1605. 给定行和列的和求可行矩阵
func restoreMatrix(rowSum []int, colSum []int) [][]int {
	m := len(rowSum)
	n := len(colSum)
	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			minVal := min(rowSum[i], colSum[j])
			arr[i][j] = minVal
			rowSum[i] -= minVal
			colSum[j] -= minVal
			if rowSum[i] == 0 && colSum[j] == 0 {
				break
			}
		}
	}
	return arr
}
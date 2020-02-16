package main

import "fmt"

func main(){
	b := []int{1}
	c := []int{2,3,4}
	a := findMedianSortedArrays(b,c)
	fmt.Println("a = ",a)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var a  float64
	if len(nums1) <= len(nums2) {
		a =H(nums1,nums2)

	}else if len(nums2) <= len(nums1) {
		a = H(nums2,nums1)

	}
	return  a
}

func H(nums1 []int, nums2 []int) float64 {
	lens := len(nums1) + len(nums2)
	f1 := 0
	f2 := 0
	index := 0
	Fuck := [10]int{}
	k := lens / 2
	var r float64
	fmt.Println("k = ",k)
	if k == 0 {
		if len(nums1) == 1 {
			return float64(nums1[0])
		}else if len(nums2) == 1 {
			return float64(nums2[0])
		}else {
			return 0
		}
	}

	if len(nums1) == 1 && len(nums2) == 1 {
		r = float64(nums1[f1] + nums2[f2]) /float64(2.0)
		return r
	}else if len(nums1) == 1 && len(nums2) > 1 {
		for {
			fmt.Println("index = ",index,"f1",f1,"f2",f2,"Fuck",Fuck)
			if  f1 >=  len(nums1) || f2 >= len(nums2)  || index == k + 1{
				fmt.Println(f1,f2)
				a := L(f1,f2,nums1,nums2)
				fmt.Println(a)
				Fuck[index] = a
				break
			}
			if nums1[f1] > nums2[f2] {
				Fuck[index] = nums2[f2]
				f2++
				index++
			}else if nums1[f1] == nums2[f2] {
				Fuck[index] = nums2[f2]
				Fuck[index+1] = nums1[f1]
				f2++
				index = index + 2
			}else if nums1[f1] < nums2[f2] {

				index++
			}
		}
		fmt.Println(Fuck)
		if lens % 2 == 0 {
			r = float64(Fuck[k]+Fuck[k - 1]) / float64(2)
		}else {
			r = float64(Fuck[k])
		}
	}else if len(nums2) == 1 && len(nums1) > 1 {

	}



	if len(nums1) == 0 && len(nums2) > 1 {
		fmt.Println("1")
		nums1s := [1]int{}
		for {
			if   f2 >= len(nums2)  || index == k {
				a := L(f1,f2,nums1,nums2)
				Fuck[index] = a
				break
			}
			if nums1s[f1] > nums2[f2] {
				Fuck[index] = nums2[f2]
				f2++
				index++
			}else if nums1s[f1] == nums2[f2] {
				Fuck[index] = nums2[f2]
				f2++
				index++
			}else if nums1s[f1] < nums2[f2] {
				Fuck[index] = nums2[f2]
				f2++
				index++
			}
		}
		fmt.Println(Fuck)
		if lens % 2 == 0 {
			r = float64(Fuck[k]+Fuck[k - 1]) / float64(2)
		}else {
			r = float64(Fuck[k])
		}
	}else if len(nums2) == 0 && len(nums1) > 1 {
		fmt.Println("2")
		nums1s := [1]int{}
		for {
			if   f2 >= len(nums1)  || index == k {
				a := L(f1,f2,nums1,nums2)
				Fuck[index] = a
				break
			}
			if nums1s[f1] > nums1[f2] {
				Fuck[index] = nums1[f2]
				f2++
				index++
			}else if nums1s[f1] == nums1[f2] {
				Fuck[index] = nums1[f2]
				f2++
				index++
			}else if nums1s[f1] < nums1[f2] {
				Fuck[index] = nums1[f2]
				f2++
				index++
			}
		}
		fmt.Println(Fuck)
		if lens % 2 == 0 {
			r = float64(Fuck[k]+Fuck[k - 1]) / float64(2)
		}else {
			r = float64(Fuck[k])
		}
	}else if lens % 2 == 0 {
		fmt.Println("3")
		for {
			fmt.Println("index = ",index,"f1",f1,"f2",f2,"Fuck",Fuck)
			if  f1 >=  len(nums1) || f2 >= len(nums2)  || index == k + 1{
				fmt.Println(f1,f2)
				a := L(f1,f2,nums1,nums2)
				fmt.Println(a)
				Fuck[index] = a
				break
			}
			if nums1[f1] > nums2[f2] {
				Fuck[index] = nums2[f2]
				f2++
				index++
			}else if nums1[f1] == nums2[f2] {
				Fuck[index] = nums2[f2]
				Fuck[index+1] = nums1[f1]
				f1++
				f2++
				index = index + 2
			}else if nums1[f1] < nums2[f2] {
				Fuck[index] = nums1[f1]
				f1++
				index++
			}
		}
		fmt.Println(Fuck)
		if lens % 2 == 0 {
			r = float64(Fuck[k]+Fuck[k - 1]) / float64(2)
		}else {
			r = float64(Fuck[k])
		}
	}else {
		fmt.Println("4")
		for {
			fmt.Println("index = ",index,"f1",f1,"f2",f2,"Fuck",Fuck)
			if  f1 >=  len(nums1) || f2 >= len(nums2)  || index == k + 1{
				a := L(f1,f2,nums1,nums2)
				Fuck[index] = a
				break
			}
			if nums1[f1] > nums2[f2] {
				Fuck[index] = nums2[f2]
				f2++
				index++
			}else if nums1[f1] == nums2[f2] {
				Fuck[index] = nums2[f2]
				Fuck[index+1] = nums1[f1]
				f1++
				f2++
				index = index + 2
			}else if nums1[f1] < nums2[f2] {
				Fuck[index] = nums1[f1]
				f1++
				index++
			}
		}
		fmt.Println(Fuck)
		if lens % 2 == 0 {
			r = float64(Fuck[k]+Fuck[k + 1]) / float64(2)
		}else {
			r = float64(Fuck[k])
		}
	}
	return r
}

func L(f1,f2 int,nums1,nums2 []int) int {
	var a int
	if f1 >= len(nums1) &&  f2 >= len(nums2) {
		a = nums2[f2-1]
	}else if f1 >= len(nums1) {
		a = nums2[f2]
		fmt.Println("L",a)
	}else if f2 >= len(nums2) {
		a = nums1[f1]
		fmt.Println("L",a)
	}
	return a
}

func R(nums []int,f2 int)  {

}
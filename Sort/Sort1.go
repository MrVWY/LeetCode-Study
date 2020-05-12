package main

//1095. 山脉数组中查找目标值  二分查找
/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
	var left, right int
	if mountainArr.length() < 3 {
		return -1
	}
	left, right = 0, mountainArr.length()-1
	for left != right {
		mid := (left + right) / 2
		if mountainArr.get(mid) < mountainArr.get(mid+1) {
			left = mid+1
		}else {
			right = mid
		}
	}
	peak := left

	left, right = 0, peak
	for left != right {
		mid := (left + right) / 2
		value := mountainArr.get(mid)
		if target == value {
			return mid
		}
		if target < value {
			right = mid
		}else {
			left = mid+1
		}
	}
	if target == mountainArr.get(left) {return left}

	left, right = peak, mountainArr.length()-1
	for left != right {
		mid := (left+right) / 2
		value := mountainArr.get(mid)
		if target == value {
			return mid
		}
		if target < value {
			left = mid + 1
		}else {
			right = mid
		}
	}
	if target == mountainArr.get(left) {return left}
	return -1
}

//33. 搜索旋转排序数组  二分查找
func search(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if (nums[left] <= target && target <= nums[mid]) || (nums[mid] <= nums[right] && (target < nums[mid] || target > nums[right]))  {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
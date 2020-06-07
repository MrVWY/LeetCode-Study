package Daily

import "sort"

//5428. 重新排列数组
func shuffle(nums []int, n int) []int {
	length := len(nums)
	mid := length/2
	start := 0
	res := make([]int, 0)
	for i := 0 ; i < length ; i++ {
		if i % 2 == 0 || i == 0 {
			res = append(res, nums[start])
			start++
		}else {
			res = append(res, nums[mid])
			mid++
		}
	}
	return res
}

//5429. 数组中的 k 个最强值
func getStrongest(arr []int, k int) []int {
	var ans = make([]int, k)
	sort.Ints(arr)
	n := len(arr)
	mid := arr[(n-1)/2]
	left, right := 0, n - 1
	for i := 0; i < k; i++ {
		hdiff, ldiff := arr[right]-mid, mid-arr[left]
		if hdiff >= ldiff {
			ans[i] = arr[right]
			right--
		} else {
			ans[i] = arr[left]
			left++
		}
	}
	return ans
}

//5430. 设计浏览器历史记录
type BrowserHistory struct {
	url []string
	cache int
	index int
}

func Constructor(homepage string) BrowserHistory {
	url := make([]string,0)
	url = append(url, homepage)
	return BrowserHistory{
		url: url,
		cache : 0,
		index: 0,
	}
}

func (this *BrowserHistory) Visit(url string)  {
	if this.index == len(this.url)-1 {
		this.url = append(this.url, url)
		this.cache++
		this.index++
	}else {
		num := len(this.url[this.index+1:])
		this.cache -= num
		this.url = this.url[:this.index+1]
		this.url = append(this.url, url)
		this.cache++
		this.index = this.cache
	}
}

func (this *BrowserHistory) Back(steps int) string {
	if this.index - steps >= 0 {
		this.index = this.index - steps
	}else {
		this.index = 0
	}
	return this.url[this.index]
}

func (this *BrowserHistory) Forward(steps int) string {
	if this.index + steps >= this.cache {
		this.index = this.cache
	}else {
		this.index = this.index + steps
	}
	return this.url[this.index]
}
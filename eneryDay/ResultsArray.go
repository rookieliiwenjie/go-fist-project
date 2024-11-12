package eneryDay

func resultsArray(nums []int, k int) []int {
	reslut := make([]int, len(nums)-k+1)
	for i := 0; i < len(nums)-k+1; i++ {
		flag := true
		for j := i + 1; j < i+k; j++ {
			if nums[j]-nums[j-1] != 1 {
				flag = false
				break
			}
		}
		if flag {
			reslut[i] = nums[i] + k - 1
		} else {
			reslut[i] = -1
		}

	}
	return reslut

}

func resultsArray2(nums []int, k int) []int {
	reslut := make([]int, len(nums)+k-1)
	for i := 0; i < len(nums)+k-1; i++ {
		reslut[i] = -1
	}
	for i := 0; i < len(nums)+k-1; i++ {

	}
	return reslut

}

package eneryDay

func countKConstraintSubstrings(s string, k int) int {
	sum := 0
	count := len(s)
	arr := []rune(s)
	for i := 0; i < count; i++ {
		zSum := 0
		oneSum := 0
		for j := i; j < count; j++ {
			if arr[j] == '0' {
				zSum++
			}
			if arr[j] == '1' {
				oneSum++
			}
			if zSum > k && oneSum > k {
				break
			}
			sum++
		}
	}
	return sum
}

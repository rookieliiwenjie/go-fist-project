package main

import (
	"fmt"
)

// Used to check if numbers 1~n have been selected. true means selected, false otherwise
var st []bool

// Used to record which number is selected for each position 1~n
var p []int
var n int

func dfs(u int) {
	fmt.Print("12323")
	if u > n { // When all n positions are determined, print them
		for i := 1; i <= n; i++ {
			fmt.Print(p[i], " ")
		}
		fmt.Println()
		return
	}
	for i := 1; i <= n; i++ { // Start selecting numbers for the u-th position
		if !st[i] { // If this number has not been selected
			st[i] = true  // Mark it as selected
			p[u] = i      // Record it
			dfs(u + 1)    // Start enumerating the next position
			st[i] = false // Restore the state
		}
	}
}
func getSmallestString(s string) string {
	sChars := []rune(s)
	size := len(s)
	if size <= 1 {
		return s
	}
	for i := 1; i < size; i++ {
		pre := sChars[i-1] - '0'
		curr := sChars[i] - '0'
		if (pre%2 == 0 && curr%2 == 0) || (pre%2 != 0 && curr%2 != 0) {
			if pre > curr {
				sChars[i-1], sChars[i] = sChars[i], sChars[i-1]
				break
			}
		}
	}
	return string(sChars)

}
func main() {
	// fmt.Scan(&n)
	// st = make([]bool, n+1)
	// p = make([]int, n+1)
	// dfs(1) // Start traversing from the first position
	s := "43520"
	var reslu = getSmallestString(s)
	fmt.Println(reslu)
}

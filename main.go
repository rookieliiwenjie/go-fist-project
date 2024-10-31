package main

import "fmt"

const N = 10

var st [N]bool // 用来判断1~n这n个数是否被选。true代表已经被选了，false反之
var p [N]int   // 用来记录1~n个位置选的是哪个数
var n int

func dfs(u int) {
	fmt.Print("go dfs")
	if u > n { // 当n个位置都确定之后就打印
		for i := 1; i <= n; i++ {
			fmt.Print(p[i], " ")
		}
		fmt.Println()
		return
	}
	for i := 1; i <= n; i++ { // 第u个位置开始选数
		if !st[i] { // 如果这个数没有被选
			st[i] = true  // 选择这个数打上标记
			p[u] = i      // 记录
			dfs(u + 1)    // 开始枚举下一个位置
			st[i] = false // 恢复现场
		}
	}
}
func main() {
	fmt.Scan(&n)
	dfs(1) // 从第一个位置开始遍历
}

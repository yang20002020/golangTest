/*
描述
输入一个 int 型的正整数，计算出该 int 型数据在内存中存储时 1 的个数。

数据范围：保证在 32 位整型数字范围内
输入描述：

	输入一个整数（int类型）

输出描述：

	这个数转换成2进制后，输出1的个数

示例1
输入：
5

输出：
2
*/
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	tem := make(map[int]int, 1)
	//tem := make([]int, 1) 越界？？？

	var j int
	for i := 0; n != 1; i++ {
		if n%2 == 0 {
			n = n / 2
			tem[i] = 0
		} else {
			n = n / 2
			tem[i] = 1
			j++
		}

	}
	if n == 1 {
		tem[len(tem)] = 1
		j++
	}

	fmt.Println(j)
	fmt.Println(tem)

}

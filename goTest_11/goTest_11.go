/*
描述
输入一个整数，将这个整数以字符串的形式逆序输出
程序不考虑负数的情况，若数字含有0，则逆序形式也含有0，如输入为100，则输出为001

数据范围： 0 \le n \le 2^{30}-1 \0≤n≤2
30

	−1

输入描述：
输入一个int整数

输出描述：
将这个整数以字符串的形式逆序输出

示例1
输入：
1516000

输出：
0006151
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int
	fmt.Scan(&N)
	str := strconv.Itoa(N)
	str1 := strReverse(str)
	fmt.Println(str1)

}
func strReverse(str string) string {
	bt := []byte(str)
	btTem := make([]byte, len(bt))
	for i := range bt {
		btTem[len(bt)-i-1] = bt[i]
	}

	return string(btTem)
}

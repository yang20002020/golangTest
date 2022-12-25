/*
描述
接受一个只包含小写字母的字符串，然后输出该字符串反转后的字符串。（字符串长度不超过1000）

输入描述：
输入一行，为一个只包含小写字母的字符串。

输出描述：
输出该字符串反转后的字符串。

示例1
输入：
abcd

输出：
dcba
*/

package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
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

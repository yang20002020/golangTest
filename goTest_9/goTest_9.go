package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
描述
输入一个 int 型整数，按照从右向左的阅读顺序，返回一个不含重复数字的新的整数。
保证输入的整数最后一位不是 0 。

数据范围： 1 <= n <= 10^8

输入描述：
输入一个int型整数

输出描述：
按照从右向左的阅读顺序，返回一个不含重复数字的新的整数

示例1
输入：
9876673

输出：
37689
*/
func main() {
	var N int = 0
	fmt.Scan(&N) //13579
	strN := strconv.Itoa(N)
	//fmt.Println(strN)
	str1 := strReverse(strN)
	str2 := deleteRep(str1)
	fmt.Println(str2)
}

func strReverse(str string) string {
	bt := []byte(str) //字符串转换成ascii码
	//fmt.Println(bt)            //数组 [49 51 53 55 57]
	//fmt.Println(string(bt[1])) //ascii码对应的字符
	btTem := make([]byte, len(bt))
	for i := range bt {
		btTem[len(bt)-i-1] = bt[i]
	}
	btem1 := string(btTem)
	//fmt.Println(btem1) //97531
	return btem1
}

func deleteRep(str string) string {
	temp := make([]byte, 1)
	temp[0] = str[0]
	for i, v := range str {
		if strings.Contains(string(temp), string(v)) {
			continue
		} else {
			temp = append(temp, str[i])
		}

	}

	return string(temp)
}

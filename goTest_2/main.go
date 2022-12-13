package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
描述
写出一个程序，接受一个由字母、数字和空格组成的字符串，和一个字符，然后输出输入字符串中该字符的出现次数。（不区分大小写字母）

数据范围： 1 \le n \le 1000 \1≤n≤1000
输入描述：
第一行输入一个由字母、数字和空格组成的字符串，第二行输入一个字符（保证该字符不为空格）。

输出描述：
输出输入字符串中含有该字符的个数。（不区分大小写字母）
*/
func main() {
	temp := make(map[byte]int)

	input := bufio.NewScanner(os.Stdin)

	if input.Scan() {

		s := input.Text()

		input.Scan() //输入一个字符   扫描

		s1 := input.Text()
		s = strings.ToLower(s)
		s1 = strings.ToLower(s1)
		b1 := s1[0]

		for i := 0; i < len(s); i++ {
			temp[s[i]]++
		}
		fmt.Println(temp[b1])

	}
}

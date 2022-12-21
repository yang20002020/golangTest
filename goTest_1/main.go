package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
描述
计算字符串最后一个单词的长度，单词以空格隔开，字符串长度小于5000。（注：字符串末尾不以空格为结尾）
输入描述：
输入一行，代表要计算的字符串，非空，长度小于5000。

输出描述：
输出一个整数，表示输入字符串最后一个单词的长度。

示例1
输入：
hello nowcoder
复制
输出：
8
复制
说明：
最后一个单词为nowcoder，长度为8
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	//在终端输入scanner.Scan()
	if scanner.Scan() {
		b := scanner.Text()
		fmt.Println(b)
		fmt.Println(strings.LastIndex(b, " "))
		a := len(b) - strings.LastIndex(b, " ") - 1
		fmt.Println(a)
	}

}

/*
当频繁地对少量数据读写时会占用IO，造成性能问题。golang的bufio库使用缓存来一次性进行大块数据的读写，以此降低IO系统调用，提升性能。
*/

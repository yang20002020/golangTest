package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
描述
•输入一个字符串，请按长度为8拆分每个输入字符串并进行输出；

•长度不是8整数倍的字符串请在后面补数字0，空字符串不处理。
输入描述：
连续输入字符串(每个字符串长度小于等于100)

输出描述：
依次输出所有分割后的长度为8的新字符串

示例1
输入：
abc

输出：
abc00000
*/

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	s := input.Text()
	tmp := []byte(s)
	for len(tmp) >= 8 { //读取的内容超过8，每次截取8位输出
		fmt.Print(string(tmp[:8]), "\n")
		tmp = tmp[8:]
	}
	for i := len(tmp); i != 0 && i < 8; i++ {
		//当剩下的长度不超过8时，补齐至8位；如果正好分完，即len(tmp)此时等于0，则直接结束
		tmp = append(tmp, '0')
	}
	fmt.Print(string(tmp))
}

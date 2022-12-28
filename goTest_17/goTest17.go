/*
描述
开发一个坐标计算工具， A表示向左移动，D表示向右移动，W表示向上移动，S表示向下移动。从（0,0）点开始移动，从输入字符串里面读取一些坐标，并将最终输入结果输出到输出文件里面。

输入：

合法坐标为A(或者D或者W或者S) + 数字（两位以内）

坐标之间以;分隔。

非法坐标点需要进行丢弃。如AA10;  A1A;  $%$;  YAD; 等。

下面是一个简单的例子 如：

A10;S20;W10;D30;X;A1A;B10A11;;A10;

处理过程：

起点（0,0）

+   A10   =  （-10,0）

+   S20   =  (-10,-20)

+   W10  =  (-10,-10)

+   D30  =  (20,-10)

+   x    =  无效

+   A1A   =  无效

+   B10A11   =  无效

+  一个空 不影响

+   A10  =  (10,-10)

结果 （10， -10）

数据范围：每组输入的字符串长度满足 1<= n <= 10000  ，坐标保证满足 -2^31  x,y <= 2^31-1
 ，且数字部分仅含正数
输入描述：
一行字符串

输出描述：
最终坐标，以逗号分隔

示例1
输入：
A10;S20;W10;D30;X;A1A;B10A11;;A10;

输出：
10,-10
示例2
输入：
ABC;AKL;DA1;

输出：
0,0
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var str string
	fmt.Scan(&str)
	strTem := strings.Split(str, ";")
	x := 0
	y := 0
	fmt.Println(strTem)
	for _, v := range strTem {
		//fmt.Println("v[len(v)-1]:", v[len(v)-1])
		//fmt.Println("byte(v[len(v)-1]):", byte(v[len(v)-1]))
		//fmt.Println("string(v[len(v)-1]):", string(v[len(v)-1]))
		//fmt.Println("v[0]:", v[0])
		//fmt.Println("string(v[0]):", string(v[0]))
		//fmt.Println("strTem[len(strTem)-1]:", strTem[len(strTem)-1])
		//fmt.Println("strTem[len(strTem)-1][1]:", strTem[len(strTem)-1][1])
		//fmt.Println("byte('A'):", byte('A'))
		if v == "" {
			continue
		}

		switch v[0] {
		// 'A'等价byte('A')
		case 'A':
			strChar := ""
			strChar = v[1:]
			strIA, err := strconv.Atoi(strChar)
			fmt.Println(strIA)
			if err != nil {
				continue
			}
			x -= strIA
		case 'D':
			strCharD := ""
			strCharD = v[1:]
			strID, err := strconv.Atoi(strCharD)
			if err != nil {
				continue
			}
			x += strID
		case 'W':
			strCharW := ""
			strCharW = v[1:]
			strIW, err := strconv.Atoi(strCharW)
			if err != nil {
				continue
			}
			y += strIW
		case 'S':
			strCharS := ""
			strCharS = v[1:]
			strIS, err := strconv.Atoi(strCharS)
			if err != nil {
				continue
			}
			y -= strIS
		default:
			continue
		}

	}
	fmt.Println(x, ",", y)
	strr := strconv.Itoa(x) + "," + strconv.Itoa(y)
	fmt.Println(strr)
}

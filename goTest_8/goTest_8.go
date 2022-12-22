package main

import (
	"fmt"
	"sort"
)

/*
描述
数据表记录包含表索引index和数值value（int范围的正整数），请对表索引相同的记录进行合并，即将相同索引的数值进行求和运算，输出按照index值升序进行输出。

提示:
0 <= index <= 11111111
1 <= value <= 100000

输入描述：
先输入键值对的个数n（1 <= n <= 500）
接下来n行每行输入成对的index和value值，以空格隔开

输出描述：
输出合并后的键值对（多行）

示例1
输入：
4
0 1
0 2
1 2
3 4
复制
输出：
0 3
1 2
*/
func main() {
	var n int
	var keyArray []int
	temp := make(map[int]int)
	fmt.Scan(&n) //先输入键值对的个数
	for i := 0; i < n; i++ {
		var key int
		var value int
		fmt.Scan(&key, &value)
		temp[key] += value
	}
	for k := range temp {
		keyArray = append(keyArray, k)
	}
	/*
		4
		9 1
		8 2
		7 7
		6 6
		key: [9 8 7 6]

	*/
	fmt.Println("key:", keyArray) //key: [9 8 7 6]
	sort.Ints(keyArray)           //按增长顺序排序
	fmt.Println("key:", keyArray) //key: [6 7 8 9]
	for _, v := range keyArray {
		fmt.Println(v, temp[v])
	}

}

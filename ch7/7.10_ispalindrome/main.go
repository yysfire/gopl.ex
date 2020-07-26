// 《The Go Programming Language》: http://gopl.io/
// 《Go语言圣经》: https://github.com/golang-china/gopl-zh
// Copyright © 2020 yysfire. All Rights Reserved.

// 练习 7.10：
// sort.Interface类型也可以适用在其它地方。
// 编写一个IsPalindrome(s sort.Interface) bool函数表明序列s是否是回文序列，
// 换句话说反向排序不会改变这个序列。假设如果!s.Less(i, j) && !s.Less(j, i)则索引i和j上的元素相等。

package main

import (
	"fmt"
	"sort"
)

//IsPalindrome 返回序列s是否为回文序列
func IsPalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len()/2; i++ {
		if s.Less(i, s.Len()-1-i) || s.Less(s.Len()-1-i, i) {
			return false
		}
	}
	return true
}

func main() {
	s := [][]string{
		[]string{"a", "b", "c"},
		[]string{"a", "b", "b", "a"},
		[]string{"上", "海", "在", "海", "上"},
		[]string{"上"},
		[]string{},
	}
	for _, v := range s {
		fmt.Printf("IsPalindrome(%s) = %v\n", v, IsPalindrome(sort.StringSlice(v)))
	}
}

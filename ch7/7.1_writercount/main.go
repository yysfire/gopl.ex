// 《The Go Programming Language》: http://gopl.io/
// 《Go语言圣经》: https://github.com/golang-china/gopl-zh
// Copyright © 2020 yysfire. All Rights Reserved.

// 练习 7.1：
// 使用来自ByteCounter的思路，实现一个针对单词和行数的计数器。
// 你会发现bufio.ScanWords非常的有用。

package main

import (
	"bufio"
	"fmt"
	"strings"
)

//ByteCounter 实现了io.Writer，进行字节数统计
type ByteCounter int

func (b *ByteCounter) Write(p []byte) (n int, err error) {
	*b += ByteCounter(len(p))
	return len(p), nil
}

//WordCounter 实现了io.Writer，进行单词数统计
type WordCounter int

func (w *WordCounter) Write(p []byte) (n int, err error) {
	pr := strings.NewReader(string(p)) //创建一个Reader，供bufio.NewScanner使用
	scanner := bufio.NewScanner(pr)
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	*w += WordCounter(count)
	return count, scanner.Err()
}

//LineCounter 实现了io.Writer，进行行数统计
type LineCounter int

func (l *LineCounter) Write(p []byte) (n int, err error) {
	pr := strings.NewReader(string(p))
	scanner := bufio.NewScanner(pr)
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	*l += LineCounter(count)
	return count, scanner.Err()
}

func main() {
	strArray := []string{
		"Hello world\nHello you!\r",
		"中国你好\n你 好 中 国\r",
	}

	var b ByteCounter
	for _, v := range strArray {
		n, err := b.Write([]byte(v))
		fmt.Printf("has %v bytes\n", n)
		if err != nil {
			fmt.Println("err:", err)
		}
	}
	fmt.Println("Total bytes:", b)

	var w WordCounter
	for _, v := range strArray {
		n, err := w.Write([]byte(v))
		fmt.Printf("has %v words\n", n)
		if err != nil {
			fmt.Println("err:", err)
		}
	}
	fmt.Println("Total words:", w)

	var l LineCounter
	for _, v := range strArray {
		n, err := l.Write([]byte(v))
		fmt.Printf("has %v lines\n", n)
		if err != nil {
			fmt.Println("err:", err)
		}
	}
	fmt.Println("Total lines:", l)
}

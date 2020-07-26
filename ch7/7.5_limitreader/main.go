// 《The Go Programming Language》: http://gopl.io/
// 《Go语言圣经》: https://github.com/golang-china/gopl-zh
// Copyright © 2020 yysfire. All Rights Reserved.

// 练习 7.5：
// io包里面的LimitReader函数接收一个io.Reader接口类型的r和字节数n，
// 并且返回另一个从r中读取字节但是当读完n个字节后就表示读到文件结束的Reader。
// 实现这个LimitReader函数：
// func LimitReader(r io.Reader, n int64) io.Reader

package main

import (
	"fmt"
	"io"
	"strings"
)

//LimiteReaderSt 封装了一个 io.Reader，并可以指定最大读取的字节数
type LimiteReaderSt struct {
	reader   io.Reader
	maxBytes int64
}

func (l *LimiteReaderSt) Read(p []byte) (n int, err error) {
	if l.maxBytes <= 0 {
		//return 0, io.EOF
		return 0, nil
	}
	if int64(len(p)) > l.maxBytes {
		p = p[:l.maxBytes]
	}

	n, err = l.reader.Read(p)
	l.maxBytes -= int64(n) // 这一点容易忽略
	return
}

//LimitReader 接收一个io.Reader接口类型的r和字节数n，
//并且返回另一个从r中读取字节但是当读完n个字节后就表示读到文件结束的Reader
func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimiteReaderSt{r, n}
}

func main() {
	r1 := LimitReader(strings.NewReader("hello world"), 10)
	buf := make([]byte, 5)
	n, err := r1.Read(buf)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Printf("Read %v bytes: %s\n", n, buf)

	n, err = r1.Read(buf)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Printf("Read %v bytes: %s\n", n, buf)

	for i := range buf {
		buf[i] = 0
	}
	n, err = r1.Read(buf)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Printf("Read %v bytes: %s\n", n, buf)
}

// 《The Go Programming Language》: http://gopl.io/
// 《Go语言圣经》: https://github.com/golang-china/gopl-zh
// Copyright © 2020 yysfire. All Rights Reserved.

// 练习 7.2：
// 写一个带有如下函数签名的函数CountingWriter，传入一个io.Writer接口类型，
// 返回一个把原来的Writer封装在里面的新的Writer类型和一个表示新的写入字节数的int64类型指针。
// func CountingWriter(w io.Writer) (io.Writer, *int64)

package main

import (
	"fmt"
	"io"
	"os"
)

//CountingWriterSt 封装了一个io.Writer，并记录写入字节数
type CountingWriterSt struct {
	w     io.Writer
	count int64
}

func (c *CountingWriterSt) Write(p []byte) (n int, err error) {
	n, err = c.w.Write(p)
	c.count += int64(n)
	return n, err
}

//CountingWriter 传入一个io.Writer接口类型，返回一个把原来的Writer
//封装在里面的新的Writer类型和一个表示新的写入字节数的int64类型指针
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := CountingWriterSt{w, 0}
	return &cw, &(cw.count)
}

func main() {
	cw, np := CountingWriter(os.Stdout)
	fmt.Fprint(cw, "hello world\t")
	fmt.Printf("%v bytes write\n", *np)
	fmt.Fprint(cw, "hello world\t")
	fmt.Printf("%v bytes write\n", *np)
	*np = 0
	fmt.Printf("%v bytes write\n", *np)
}

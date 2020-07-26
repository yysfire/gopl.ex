// 《The Go Programming Language》: http://gopl.io/
// 《Go语言圣经》: https://github.com/golang-china/gopl-zh
// Copyright © 2020 yysfire. All Rights Reserved.

// 练习 1.8：
// 修改fetch这个范例，如果输入的url参数没有 http:// 前缀的话，
// 为这个url加上该前缀。你可能会用到strings.HasPrefix这个函数。

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			url = strings.Replace(url, "http://", "https://", 1)
			fmt.Println("try to fetch:", url)
			resp, err = http.Get(url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
				os.Exit(1)
			}
		}
		fmt.Println("response status code:", resp.StatusCode)
		defer resp.Body.Close()
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

//!-

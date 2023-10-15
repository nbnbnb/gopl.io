// 从标准输入读入文件名，然后将其转换位基本文件名（去掉.和后缀）输出
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(basename(input.Text()))
	}
}

// 另一种实现
// basename removes directory components and a trailing .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	// -1 if "/" not found
	slash := strings.LastIndex(s, "/")

	// 去除最后一个 '/' 和之前的所有字符
	s = s[slash+1:]

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		// 去除最后一个 '.' 和之后的所有字符
		s = s[:dot]
	}
	return s
}

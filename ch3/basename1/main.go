// 从标准输入读入文件名，然后将其转换位基本文件名（去掉.和后缀）输出
package main

import (
	"bufio"
	"fmt"
	"os"
)

/*

测试方式 - CMD 中执行

go run .\main.go
a.txt
回车
b.iso
回车
c.java
回车
*/

func main() {
	input := bufio.NewScanner(os.Stdin)
	// Scan每次读取一行，去掉换行符，返回 bool值
	//  如果读到了内容，返回 true，否则返回 false
	for input.Scan() {
		fmt.Println(basename(input.Text()))
	}
}

// basename removes directory components and a .suffix.
// 示例 a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			// 去掉最后一个 '/' 和之前的所有字符
			s = s[i+1:]
			break
		}
	}

	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			// 去掉最后一个 '.' 和之后的所有字符
			s = s[:i]
			break
		}
	}
	return s
}

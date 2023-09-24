// Echo4 prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"strings"
)

var (
	n   = flag.Bool("n", false, "omit trailing newline")
	sep = flag.String("s", " ", "separator")
)

func main() {
	// 当程序运行时，必须在使用标志参数对应的变量之前先调用 flag.Parse 函数
	// 用于更新每个标志参数对应变量的值（之前是默认值）
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

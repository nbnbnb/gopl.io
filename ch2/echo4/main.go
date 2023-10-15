// 解析命令行参数
package main

import (
	// 导入一个 flag 包，用于处理命令行参数
	"flag"
	"fmt"
	"strings"
)

// 设置变量的默认值
// 注意，返回的都是指针类型
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

/*
支持的 flag 类型
-flag
--flag   // double dashes are also permitted
-flag=x
-flag x  // non-boolean flags only

使用方式
go run .\main.go -n=true -s='&' a=1 b=2 c=3
go run .\main.go --n -s='&' a=1 b=2 c=3
go run .\main.go -n -s='&' a=1 b=2 c=3

go run .\main.go -n -s '&' a=1 b=2 c=3

输出
a=1&b=2&c=3
*/

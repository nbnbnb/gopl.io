package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 使用短变量声明
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	// 注意：切换到命令行模式执行（不能使用 Debug 模式）
	// Linux 下用 Ctrl+D 终止
	// Windows 下用 Ctrl+C 终止（或者 Ctrl+Z，然后在输入回车）

	// 该变量从程序的标准输入中读取内容
	// 每次调用 input.Scan()，即读入下一行，并移除行末的换行符
	for input.Scan() {
		fmt.Println("read line:", input.Text())
		counts[input.Text()]++
	}

	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}

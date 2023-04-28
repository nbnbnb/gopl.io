package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	// 注意：Linux 下用 Ctrl+D 终止
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}

package main

import "fmt"

// 声明一个常量（float）
const boilingF = 212.0

func main() {
	f := boilingF
	c := (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}

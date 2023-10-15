// 华氏度转摄氏度
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0

	// 调用外挂函数 fToC
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))
}

// 给 float64 外挂一个函数
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

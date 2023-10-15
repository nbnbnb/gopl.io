// 转换摄氏度和华氏度
package main

import (
	"fmt"
	"os"
	"strconv"

	// 导入一个包（本地包）
	"gopl.io/ch2/tempconv"
)

func main() {
	// 如果没有参数，则默认为 12.3
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "12.3")
	}

	for _, arg := range os.Args[1:] {
		// strconv 是框架自带的包
		// 转换得到一个 float64（float64 is the set of all IEEE-754 64-bit floating-point numbers.）
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		// 转换成自定义的类型
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)

		// 调用自定义 tempconv 包中的函数（大写字母开头，公共的可调用）
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}

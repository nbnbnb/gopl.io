// 一个自定义包，用于转换摄氏度和华氏度
package tempconv

import "fmt"

// 声明一个类型（多个一起）
// 大写字母开头，表示公共
type (
	Celsius    float64
	Fahrenheit float64
)

// 声明常量（多个一起）
// 大写字母开头，表示公共
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// 给类型添加函数
// 大写字母开头，表示公共
func (c Celsius) String() string    { return fmt.Sprintf("%g°", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

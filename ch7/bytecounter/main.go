package main

import (
	"fmt"
)

type ByteCounter int

// 给 ByteCounter 添加 Write 方法(使其实现 io.Writer 接口)
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	// 调用外挂的 Write 方法
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	name := "Dolly"
	// Fprintf 的第一个参数是 io.Writer 接口类型
	// 由于 ByteCounter（指针）实现了该接口，所以签名匹配
	// 注意：此处传递的是指针，因为 ByteCounter 的 Write 方法是指针接收者
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}

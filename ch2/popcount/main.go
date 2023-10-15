package popcount

// 一个 256 字节的数组
var pc [256]byte

// 包的初始函数
func init() {
	// 初始化 pc
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount 算法
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

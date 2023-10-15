// 包文件以 _test 结尾，表明是一个测试包
package popcount_test

import (
	// 导入 testing 包
	"testing"

	// 导入 popcount 包（main.go 里面会执行初始化）
	"gopl.io/ch2/popcount"
)

// -- Alternative implementations --

func BitCount(x uint64) int {
	// Hacker's Delight, Figure 5-2.
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}

func PopCountByClearing(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		n++
	}
	return n
}

func PopCountByShifting(x uint64) int {
	n := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			n++
		}
	}
	return n
}

// -- Benchmarks --

// Benchmark
// 所有方法以 Benchmark* 开头，参数为 *testing.B

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 使用 PopCount
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkBitCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 使用 BitCount
		BitCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 使用 PopCountByClearing
		PopCountByClearing(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 使用 PopCountByShifting
		PopCountByShifting(0x1234567890ABCDEF)
	}
}

// 注意：在 CMD 中运行，PowerShell 中不输出 Benchmark 结果
/*
	d:\GitRepo\gopl.io\ch2\popcount>go test -cpu=16 -bench=. gopl.io/ch2/popcount
	goos: windows
	goarch: amd64
	pkg: gopl.io/ch2/popcount
	cpu: AMD Ryzen 7 5800H with Radeon Graphics
	BenchmarkPopCount-16                    1000000000               0.2284 ns/op
	BenchmarkBitCount-16                    1000000000               0.2303 ns/op
	BenchmarkPopCountByClearing-16          141264786                8.454 ns/op
	BenchmarkPopCountByShifting-16          77231010                15.77 ns/op
	PASS
	ok      gopl.io/ch2/popcount    4.036s
*/

// ----------------------------------------------------------------------------

// Go 1.6, 2.67GHz Xeon
// $ go test -cpu=4 -bench=. gopl.io/ch2/popcount
// BenchmarkPopCount-4                  200000000         6.30 ns/op
// BenchmarkBitCount-4                  300000000         4.15 ns/op
// BenchmarkPopCountByClearing-4        30000000         45.2 ns/op
// BenchmarkPopCountByShifting-4        10000000        153 ns/op
//
// Go 1.6, 2.5GHz Intel Core i5
// $ go test -cpu=4 -bench=. gopl.io/ch2/popcount
// BenchmarkPopCount-4                  200000000         7.52 ns/op
// BenchmarkBitCount-4                  500000000         3.36 ns/op
// BenchmarkPopCountByClearing-4        50000000         34.3 ns/op
// BenchmarkPopCountByShifting-4        20000000        108 ns/op
//
// Go 1.7, 3.5GHz Xeon
// $ go test -cpu=4 -bench=. gopl.io/ch2/popcount
// BenchmarkPopCount-12                 2000000000        0.28 ns/op
// BenchmarkBitCount-12                 2000000000        0.27 ns/op
// BenchmarkPopCountByClearing-12       100000000        18.5 ns/op
// BenchmarkPopCountByShifting-12       20000000         70.1 ns/op

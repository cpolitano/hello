package fundamentals

import "fmt"

const (
	_ = iota // 0
	KB = 1 << (iota * 10) // 1 << 1 * 10
	MB = 1 << (iota * 10) // 1 << 2 * 10
	GB = 1 << (iota * 10) // 1 << 3 * 10
)

func Constants() {
	fmt.Printf("binary\t\tdecimal\n")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
}
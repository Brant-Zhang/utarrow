package main

import "fmt"

type bits []byte

const length = 1024

func newBitMap() bits {
	b := make(bits, length<<3)
	return b
}

func (b bits) Input(data []int) {
	for _, v := range data {
		b[v] = 1
	}
}
func (b bits) Sort() []int {
	var re = make([]int, 0)
	for k, v := range b {
		if v == 1 {
			re = append(re, k)
		}
	}
	return re
}

func main() {
	x := newBitMap()
	x.Input([]int{234, 188, 999, 23, 78, 66, 2})
	fmt.Println(x.Sort())
}

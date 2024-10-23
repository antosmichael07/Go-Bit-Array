package main

import "fmt"

type BitArray struct {
	Array []byte
	Len   int
}

func NewBitArray(n int) BitArray {
	if n%8 == 0 {
		return BitArray{Array: make([]byte, n/8), Len: n}
	} else {
		return BitArray{Array: make([]byte, n/8+1), Len: n&-8 + 8}
	}
}

func (barr *BitArray) On(i int) {
	barr.Array[i/8] |= 1 << (i % 8)
}

func (barr *BitArray) Off(i int) {
	if barr.Get(i) {
		barr.Array[i/8] ^= 1 << (i % 8)
	}
}

func (barr *BitArray) Get(i int) bool {
	return barr.Array[i/8]&(1<<(i%8)) != 0
}

func main() {
	for i := 1; i <= 128; i++ {
		barr := NewBitArray(i)

		fmt.Printf("%3d, %3d\n", i, barr.Len)
	}
}

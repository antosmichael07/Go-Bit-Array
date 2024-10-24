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
		return BitArray{Array: make([]byte, n/8+1), Len: n}
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

func Slice(barr BitArray, start int, end int) BitArray {
	new_barr := NewBitArray(end - start)

	for i := start; i < end; i++ {
		if barr.Get(i) {
			new_barr.On(i - start)
		}
	}

	return new_barr
}

func Append(barr1 BitArray, barr2 BitArray) BitArray {
	new_barr := NewBitArray(barr1.Len + barr2.Len)

	for i := 0; i < barr1.Len; i++ {
		if barr1.Get(i) {
			new_barr.On(i)
		}
	}

	for i := 0; i < barr2.Len; i++ {
		if barr2.Get(i) {
			new_barr.On(barr1.Len + i)
		}
	}

	return new_barr
}

func (barr *BitArray) Slice(start int, end int) {
	new_barr := NewBitArray(end - start)

	for i := start; i < end; i++ {
		if barr.Get(i) {
			new_barr.On(i - start)
		}
	}

	*barr = new_barr
}

func (barr1 *BitArray) Append(barr2 BitArray) {
	new_barr := NewBitArray(barr1.Len + barr2.Len)

	for i := 0; i < barr1.Len; i++ {
		if barr1.Get(i) {
			new_barr.On(i)
		}
	}

	for i := 0; i < barr2.Len; i++ {
		if barr2.Get(i) {
			new_barr.On(barr1.Len + i)
		}
	}

	*barr1 = new_barr
}

func main() {
	barr1 := NewBitArray(5)
	barr1.On(1)
	barr1.On(3)
	barr2 := NewBitArray(5)
	barr2.On(2)

	barr3 := Slice(barr1, 1, 4)

	for i := 0; i < barr3.Len; i++ {
		fmt.Println(barr3.Get(i))
	}
}

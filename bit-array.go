package main

type BitArray struct {
	Array []byte
	Len   int
}

func makeBytes(n int) []byte {
	if n%8 == 0 {
		return make([]byte, n/8)
	} else {
		return make([]byte, n/8+1)
	}
}

func NewBitArray(n int) BitArray {
	return BitArray{Array: makeBytes(n), Len: n}
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

func (barr *BitArray) Slice(start int, end int) BitArray {
	new_barr := NewBitArray(end - start)

	for i := start; i < end; i++ {
		if barr.Get(i) {
			new_barr.On(i - start)
		}
	}

	return new_barr
}

func (barr *BitArray) SliceSet(start int, end int) {
	new_barr := NewBitArray(end - start)

	for i := start; i < end; i++ {
		if barr.Get(i) {
			new_barr.On(i - start)
		}
	}

	*barr = new_barr
}

func Append(barr1 *BitArray, barr2 *BitArray) BitArray {
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

func (barr1 *BitArray) AppendSet(barr2 *BitArray) {
	old_barr1 := *barr1
	barr1.Array = makeBytes(barr1.Len + barr2.Len)
	barr1.Len += barr2.Len

	for i := 0; i < old_barr1.Len; i++ {
		if old_barr1.Get(i) {
			barr1.On(i)
		}
	}

	for i := 0; i < barr2.Len; i++ {
		if barr2.Get(i) {
			barr1.On(old_barr1.Len + i)
		}
	}
}

func (barr *BitArray) SetAllOn() {
	for i := 0; i < barr.Len; i++ {
		barr.On(i)
	}
}

func (barr *BitArray) SetAllOff() {
	for i := 0; i < barr.Len; i++ {
		barr.Off(i)
	}
}

func (barr *BitArray) Resize(n int) {
	old_barr := *barr
	barr.Array = makeBytes(n)
	barr.Len = n

	for i := 0; i < min(n, old_barr.Len); i++ {
		if old_barr.Get(i) {
			barr.On(i)
		}
	}
}

func main() {
	barr1 := NewBitArray(5)
	barr1.On(1)
	barr1.On(3)
	barr2 := NewBitArray(5)
	barr2.On(2)

	barr1.Resize(8)

	for i := 0; i < barr1.Len; i++ {
		println(barr1.Get(i))
	}
}

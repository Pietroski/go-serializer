package bytesx

type Writer struct {
	data   []byte
	cursor int

	freeCap int // cap(data) - len(data)
}

func NewWriter(data []byte) *Writer {
	capacity := cap(data)
	length := len(data)
	bbw := &Writer{
		data:    data,
		freeCap: capacity - length,
	}
	if bbw.freeCap == 0 {
		bbw.freeCap = capacity
	}
	if length == 0 {
		bbw.data = bbw.data[:capacity]
	}

	return bbw
}

func (bbw *Writer) Put(b byte) {
	if 1 >= bbw.freeCap {
		newDataCap := cap(bbw.data) << 1
		newData := make([]byte, newDataCap)
		copy(newData, bbw.data)
		bbw.data = newData
		bbw.freeCap = newDataCap - bbw.cursor
	}

	bbw.data[bbw.cursor] = b
	bbw.cursor++
	bbw.freeCap--
}

func (bbw *Writer) Write(bs []byte) {
	bsLen := len(bs)
	if bsLen > bbw.freeCap {
		newCap := cap(bbw.data) << 1
		currentMaxSize := len(bbw.data) + bsLen - bbw.freeCap
		for currentMaxSize > newCap {
			newCap <<= 1
		}

		newData := make([]byte, newCap)
		copy(newData, bbw.data)
		bbw.data = newData
		bbw.freeCap = newCap - bbw.cursor
	}

	copy(bbw.data[bbw.cursor:], bs)
	bbw.cursor += bsLen
	bbw.freeCap -= bsLen
}

func (bbw *Writer) Bytes() []byte {
	return bbw.data[:bbw.cursor]
}

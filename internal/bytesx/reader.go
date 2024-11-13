package bytesx

type Reader struct {
	data   []byte
	cursor int
}

func NewReader(data []byte) *Reader {
	return &Reader{
		data: data,
	}
}

func (bbr *Reader) Next() byte {
	bbr.cursor++
	return bbr.data[bbr.cursor-1]
}

func (bbr *Reader) Read(n int) []byte {
	bbr.cursor += n
	return bbr.data[bbr.cursor-n : bbr.cursor]
}

func (bbr *Reader) Yield() int {
	return bbr.cursor
}

func (bbr *Reader) Skip(n int) {
	bbr.cursor += n
}

func (bbr *Reader) Bytes() []byte {
	return bbr.data[:bbr.cursor]
}

func (bbr *Reader) BytesFromCursor() []byte {
	return bbr.data[bbr.cursor:]
}

func (bbr *Reader) CutBytes() []byte {
	bbr.data = bbr.data[bbr.cursor:]
	bbr.cursor = 0
	return bbr.data
}

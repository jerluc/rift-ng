package rc

import (
	"bufio"
	"encoding/binary"
	"io"
)

type rcReader struct{
	reader *bufio.Reader
}

var byteOrder = binary.BigEndian

func NewRCReader(reader io.Reader) *rcReader {
	return &rcReader{bufio.NewReader(reader)}
}

// Reads the next variable number of bytes
// from the internal io.Reader, deserializing
// its contents into the provided reference.
//
// Note that this assumes the input bytes are
// serialized in Little Endian.
func (r *rcReader) Read(v interface{}) error {
	err := binary.Read(r.reader, byteOrder, v)
	return err
}

func (r *rcReader) ReadByte() byte {
	var n byte
	r.Read(&n)
	return n
}

func (r *rcReader) ReadBytes(n int) []byte {
	b := make([]byte, n)
	r.reader.Read(b)
	return b
}

func (r *rcReader) ReadUInt8() uint8 {
	var n uint8
	r.Read(&n)
	return n
}

func (r *rcReader) ReadUInt16() uint16 {
	var n uint16
	r.Read(&n)
	return n
}

func (r *rcReader) ReadUInt32() uint32 {
	var n uint32
	r.Read(&n)
	return n
}

// TODO: Is this correct?
func (r *rcReader) Finished() bool {
	return r.reader.Buffered() == 0
}


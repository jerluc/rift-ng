package rc

import (
	"encoding/binary"
	"io"
)

type binaryReader struct{
	reader io.Reader
}

var byteOrder = binary.LittleEndian

// Reads the next variable number of bytes
// from the internal io.Reader, deserializing
// its contents into the provided reference.
//
// Note that this assumes the input bytes are
// serialized in Little Endian.
func (r *binaryReader) Read(v interface{}) error {
	err := binary.Read(r.reader, byteOrder, v)
	return err
}

func (r *binaryReader) ReadByte() byte {
	var n byte
	r.Read(&n)
	return n
}

func (r *binaryReader) ReadBytes(n int) []byte {
	b := make([]byte, n)
	r.reader.Read(b)
	return b
}

func (r *binaryReader) ReadUInt8() uint8 {
	var n uint8
	r.Read(&n)
	return n
}

func (r *binaryReader) ReadUInt16() uint16 {
	var n uint16
	r.Read(&n)
	return n
}

func (r *binaryReader) ReadUInt32() uint32 {
	var n uint32
	r.Read(&n)
	return n
}

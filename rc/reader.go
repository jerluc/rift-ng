package rc

import (
	"encoding/binary"
	"io"
)

type binaryReader struct{
	reader io.Reader
}

var byteOrder = binary.LittleEndian

func (r *binaryReader) Read(v interface{}) error {
	err := binary.Read(r.reader, byteOrder, v)
	return err
}

func (r *binaryReader) ReadByte() byte {
	var n byte
	r.Read(&n)
	return n
}

func (r *binaryReader) ReadUInt32() uint32 {
	var n uint32
	r.Read(&n)
	return n
}
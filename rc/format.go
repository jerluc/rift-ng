package rc

import (
	"fmt"
	"io"
)

const CurrentVersion = 1

type RDef struct{
	rcVersion uint8
	defName   string
	defType   uint8
	data      []byte
}

// Loads and parses an RC stream per the RC
// specification (TBD).
//
// Note that RC streams are generally meant to be
// backward-compatible, though this may be
// subject to change!
//
// TODO: It would be nice to use some kind of
// binary-based PEG grammar for parsing and
// interpreting RC streams instead of hand-coding
// each element of the stream, as things such as
// run-lengths and byte offsets are pretty silly
// and honestly only make sense when it comes to
// using the data stream as a static binary blob
// but makes way less sense when it comes to
// reading arbitrary streams off of a TCP socket
func LoadStream(stream io.Reader) (*RDef, error) {
	reader := binaryReader{stream}

	// Check for the Rift header
	if head := reader.ReadByte(); head != 'R' {
		return nil, fmt.Errorf("Missing RDef head")
	}

	// Ensure that this is a parseable version
	rcVersion := reader.ReadUInt8()
	if rcVersion > CurrentVersion {
		return nil, fmt.Errorf("RC version %d is not supported", rcVersion)
	}

	// Parse the RDef
	defNameLength := int(reader.ReadUInt8())
	if defNameLength == 0 {
		return nil, fmt.Errorf("Def name must be one or more characters long")
	}

	defName := string(reader.ReadBytes(defNameLength))
	defType := reader.ReadUInt8()
	dataLength := reader.ReadUInt16()
	data := reader.ReadBytes(int(dataLength))

	if tail := reader.ReadByte(); tail != 'R' {
		return nil, fmt.Errorf("Missing RDef tail")
	}

	return &RDef{
		rcVersion,
		defName,
		defType,
		data,
	}, nil
}

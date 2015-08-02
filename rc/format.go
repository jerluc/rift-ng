package rc

import (
	"fmt"
	"io"
)

const CurrentVersion = 1

type RDef struct{
	Version uint8
	Name    string
	Type    uint8
	Code    []byte
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
	reader := NewRCReader(stream)

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
	codeLength := reader.ReadUInt16()
	code := reader.ReadBytes(int(codeLength))

	if !reader.Finished() {
		return nil, fmt.Errorf("Corrupt RDef code section")
	}

	return &RDef{
		rcVersion,
		defName,
		defType,
		code,
	}, nil
}

package rc

import (
	"fmt"
	"os"
)

type RCFile struct{}

func LoadFile(rcFilePath string) (*RCFile, error) {
	file, openErr := os.Open(rcFilePath)
	if openErr != nil {
		return nil, openErr
	}

	fi, _ := file.Stat()

	reader := binaryReader{file}
	// head := reader.ReadByte()
	// if head != 'R' {
	// 	panic("FUCKKK")
	// }

	for i := 0; i < int(fi.Size()); i++ {
		b := reader.ReadByte()
		fmt.Printf("%c\n", b)
	}

	return nil, nil
}
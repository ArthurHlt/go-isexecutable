package isexecutable

import (
	"bytes"
	"debug/elf"
	"debug/macho"
	"encoding/binary"
	"io"
)

func chunk(reader io.Reader, chunkSize int64, closeAfterCheck bool) ([]byte, error) {
	buf := new(bytes.Buffer)
	_, err := io.CopyN(buf, reader, chunkSize)
	if err != nil && err != io.EOF {
		return []byte{}, err
	}
	if !closeAfterCheck {
		return buf.Bytes(), nil
	}
	// we don't want to retrieve everything
	// so we close if it closeable
	if closer, ok := reader.(io.Closer); ok {
		err := closer.Close()
		if err != nil {
			return []byte{}, err
		}
	}
	return buf.Bytes(), nil
}

// check if source is an executable file
func IsExecutable(reader io.Reader, closeAfterCheck bool) bool {
	buf, err := chunk(reader, 4, closeAfterCheck)
	if err != nil {
		panic(err)
	}
	if len(buf) < 4 {
		return false
	}
	le := binary.LittleEndian.Uint32(buf)
	be := binary.BigEndian.Uint32(buf)

	return string(buf) == elf.ELFMAG || // elf - linux format exec file
	// .exe windows
		string(buf[:2]) == "MZ" ||
	// shebang
		string(buf[:2]) == "#!" ||
	// mach-o - mac format exec file
		macho.Magic32 == le || macho.Magic32 == be || macho.Magic64 == le || macho.Magic64 == be || macho.MagicFat == le || macho.MagicFat == be
}

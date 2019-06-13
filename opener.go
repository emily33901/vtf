package vtf

import (
	"io"
	"os"
)

// ReadFromStream loads a vtf from standard
// io.Reader stream
func ReadFromStream(stream io.Reader) (*Vtf, error) {
	reader := &Reader{
		stream: stream,
	}

	return reader.Read()
}

// ReadHeaderFromStream only reads the header from a vtf
func ReadHeaderFromStream(stream io.Reader) (*Header, error) {
	reader := &Reader{
		stream: stream,
	}

	return reader.ReadHeader()
}

// ReadFromFile is a wrapper for ReadFromStream wrapper to load directly from
// filesystem. Exists for convenience
func ReadFromFile(filepath string) (*Vtf, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	v, err := ReadFromStream(file)
	return v, err
}

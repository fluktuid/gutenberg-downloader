package util

import (
	"compress/gzip"
	"io"
)

func UnzipData(reader io.Reader, writer io.Writer) error {
	r, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}

	io.Copy(writer, r)
	return nil
}

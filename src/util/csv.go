package util

import (
	"encoding/csv"
	"io"
)

func GetCSVRecords(r io.Reader, out chan<- []string) error {
	rdr := csv.NewReader(r)
	record, err := rdr.Read()
	for record != nil {
		out <- record
		record, err = rdr.Read()
	}
	if err != io.EOF {
		return err
	}
	return nil
}

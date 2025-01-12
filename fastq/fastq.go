package fastq

import (
	"bufio"
	"fmt"
	"io"
)

type FastqRecord struct {
	Name string
	Read string
}

var ScanError = fmt.Errorf("Premature end of file inside record")

func scanRecord(s *bufio.Scanner) (*FastqRecord, error) {
	name := s.Text()[1:]
	if !s.Scan() {
		return nil, ScanError
	}
	read := s.Text()

	return &FastqRecord{name, read}, nil
}

type ReadCallback func(*FastqRecord)

func ScanFastq(r io.Reader, fn ReadCallback) error {
	s := bufio.NewScanner(r)

	for s.Scan() {
		if rec, err := scanRecord(s); err != nil {
			return err
		} else {
			fn(rec)
		}
	}

	return nil
}

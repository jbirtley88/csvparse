package csvparse

import (
	"encoding/csv"
	"io"
)

type DictReader struct {
	reader    *csv.Reader
	headerRow []string
}

func (r *DictReader) Read() (map[string]string, error) {
	var err error

	if r.headerRow == nil {
		r.headerRow, err = r.reader.Read()

		if err != nil {
			return nil, err
		}
	}

	row, err := r.reader.Read()

	if err != nil {
		return nil, err
	}

	dict := map[string]string{}

	for i, col := range r.headerRow {
		dict[col] = row[i]
	}

	return dict, nil
}

func (r *DictReader) ReadAll() ([]map[string]string, error) {
	records := make([]map[string]string, 0)

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		records = append(records, record)
	}

	return records, nil
}

func NewDictReader(r *csv.Reader) *DictReader {
	return &DictReader{
		reader: r,
	}
}

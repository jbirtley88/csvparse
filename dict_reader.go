package csvparse

import (
	"encoding/csv"
	"io"
)

// DictReader allows reading from csv.Reader for CSV files that have a header row.
// It will map each row to a map[string]string using the headings as the map keys.
type DictReader struct {
	reader    *csv.Reader
	headerRow []string
}

func mapRow(headers []string, row []string) (map[string]any, error) {
	if len(headers) != len(row) {
		return nil, ErrHeaderCountDoesNotMatchRowCount
	}

	dict := map[string]any{}

	for i, col := range headers {
		dict[col] = row[i]
	}

	return dict, nil
}

// Headers returns the header row.
func (r *DictReader) Headers() ([]string, error) {
	var err error

	if r.headerRow == nil {
		r.headerRow, err = r.reader.Read()

		if err != nil {
			return nil, err
		}
	}

	return r.headerRow, nil
}

// Read returns the next line from the *csv.Reader as a map.
// Just like csv.Reader.Read() it will return an io.EOF if no more lines are found.
func (r *DictReader) Read() (map[string]any, error) {
	headers, err := r.Headers()

	if err != nil {
		return nil, err
	}

	row, err := r.reader.Read()

	if err != nil {
		return nil, err
	}

	dict, err := mapRow(headers, row)

	if err != nil {
		return nil, err
	}

	return dict, nil
}

// Read returns the next line from the *csv.Reader as a slice of maps.
func (r *DictReader) ReadAll() ([]map[string]any, error) {
	headers, err := r.Headers()

	if err != nil {
		return nil, err
	}

	records := make([]map[string]any, 0)

	for {
		row, err := r.reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		dict, err := mapRow(headers, row)

		if err != nil {
			return nil, err
		}

		records = append(records, dict)
	}

	return records, nil
}

func NewDictReader(r *csv.Reader) *DictReader {
	return &DictReader{
		reader: r,
	}
}

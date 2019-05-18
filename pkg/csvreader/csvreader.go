package csvreader

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"

	"github.com/ray1729/goml/pkg/dataset"
)

type ValCoercer func(string) (int, error)

func NewEnumerator(xs ...string) ValCoercer {
	return func(s string) (int, error) {
		s = strings.TrimSpace(s)
		for i, x := range xs {
			if s == x {
				return i, nil
			}
		}
		return -1, fmt.Errorf("Failed to enumerate %s", s)
	}
}

type RowCoercer func([]string, []string) ([]int, error)

func NewRowCoercer(spec map[string]ValCoercer) RowCoercer {
	return func(xs, header []string) ([]int, error) {
		var result []int
		if len(xs) != len(header) {
			return result, fmt.Errorf("Header/row length mismatch")
		}
		for i, col := range header {
			f, ok := spec[col]
			if !ok {
				// Silently ignore columns with no coercer
				continue
			}
			v, err := f(xs[i])
			if err != nil {
				return result, err
			}
			result = append(result, v)
		}
		return result, nil
	}
}

func ReadCSV(r io.Reader, coerce RowCoercer) (*dataset.Dataset, error) {
	rdr := csv.NewReader(r)
	header, err := rdr.Read()
	if err != nil {
		return nil, err
	}
	ds := new(dataset.Dataset)
	ds.Header = header
	for {
		row, err := rdr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		data, err := coerce(row, header)
		if err != nil {
			return nil, err
		}
		ds.Data = append(ds.Data, data)
	}
	return ds, nil
}

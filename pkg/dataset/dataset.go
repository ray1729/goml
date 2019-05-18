package dataset

import "fmt"

type Dataset struct {
	Header []string
	Data   [][]int
}

func (ds *Dataset) ColIndex(colname string) int {
	for i, v := range ds.Header {
		if v == colname {
			return i
		}
	}
	panic(fmt.Sprintf("Dataset has no column %s", colname))
}

func (ds *Dataset) Col(colname string) []int {
	ix := ds.ColIndex(colname)
	result := make([]int, len(ds.Data))
	for i, v := range ds.Data {
		result[i] = v[ix]
	}
	return result
}

func (ds *Dataset) PartitionBy(colname string) map[int]*Dataset {
	ix := ds.ColIndex(colname)
	parts := make(map[int]*Dataset)
	for _, row := range ds.Data {
		v := row[ix]
		if _, ok := parts[v]; !ok {
			parts[v] = &Dataset{Header: ds.Header}
		}
		parts[v].Data = append(parts[v].Data, row)
	}
	return parts
}

func (ds *Dataset) ColFreq(colname string) map[int]int {
	result := make(map[int]int)
	for _, v := range ds.Col(colname) {
		result[v]++
	}
	return result
}

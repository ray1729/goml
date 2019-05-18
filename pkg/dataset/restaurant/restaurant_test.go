package restaurant

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadDataset(t *testing.T) {
	_, err := ReadDataset()
	assert.NoError(t, err, "Failed to read dataset")
}

func TestPartitionBy(t *testing.T) {
	ds, err := ReadDataset()
	if assert.NoError(t, err) {
		xs := ds.PartitionBy("Wait")
		assert.Equal(t, 2, len(xs), "Wait has 2 distinct values")
		ys := ds.PartitionBy("Price")
		assert.Equal(t, 3, len(ys), "Price has 3 distinct values")
	}
}

package decisiontree

import (
	"fmt"
	"github.com/ray1729/goml/pkg/dataset/restaurant"
	"testing"

	"github.com/stretchr/testify/assert"
)

func repeat(x, n int) []int {
	result := make([]int, n)
	for i := range result {
		result[i] = x
	}
	return result
}

func TestEntropy(t *testing.T) {
	xs := []int{0, 1}
	assert.InDelta(t, 1.0, Entropy(xs), 0.001, "entropy of fair coin")
	ys := append(repeat(0, 99), 1)
	assert.InDelta(t, 0.08, Entropy(ys), 0.001, "entropy of biassed coin")
}

func TestB(t *testing.T) {
	assert.InDelta(t, 1.0, B(0.5), 0.001, "entropy of fair coin")
	assert.InDelta(t, 0.08, B(0.99), 0.001, "entropy of biassed coin")
}

func TestGain(t *testing.T) {
	ds, err := restaurant.ReadDataset()
	if assert.NoError(t, err) {
		goal := "Wait"
		for _, attr := range ds.Header {
			if attr == goal {
				continue
			}
			g := Gain(ds, attr, goal)
			fmt.Printf("%s => %f\n", attr, g)
		}
	}
}

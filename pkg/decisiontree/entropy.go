package decisiontree

import (
	"github.com/ray1729/goml/pkg/dataset"
	"math"
)

func Entropy(observations []int) float64 {
	P := make(map[int]float64)
	for _, x := range observations {
		P[x]++
	}
	n := float64(len(observations))
	e := 0.0
	for _, p := range P {
		p = p / n
		e -= math.Log2(p) * p
	}
	return e
}

// B returns the entropy of a boolean random variable that is true with probability q
func B(q float64) float64 {
	return -(math.Log2(q)*q + math.Log2(1-q)*(1-q))
}

func Remainder(ds *dataset.Dataset, attr, goal string) float64 {
	result := 0.0
	for _, p := range ds.PartitionBy(attr) {
		result += (float64(len(p.Data)) / float64(len(ds.Data))) * Entropy(p.Col(goal))
	}
	return result
}

func Gain(ds *dataset.Dataset, attr string, goal string) float64 {
	return Entropy(ds.Col(goal)) - Remainder(ds, attr, goal)
}

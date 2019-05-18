package decisiontree

import (
	"fmt"
	"github.com/ray1729/goml/pkg/dataset/restaurant"
	"testing"
)

func TestDecisionTreeLearn(t *testing.T) {
	ds := restaurant.MustReadDataset()
	goal := ds.Header[len(ds.Header)-1]
	attributes := ds.Header[1 : len(ds.Header)-1]
	tree := DecisionTreeLearn(ds, attributes, goal, nil)
	fmt.Println(tree)
}

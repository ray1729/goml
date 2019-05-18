package decisiontree

import "github.com/ray1729/goml/pkg/dataset"

type DecisionTree struct {
	Attribute string
	IsLeaf    bool
	Decision  int
	Children  map[int]DecisionTree
}

func DecisionTreeLearn(examples *dataset.Dataset, attributes []string, goal string, parentExamples *dataset.Dataset) DecisionTree {
	result := DecisionTree{}
	if len(examples.Data) == 0 {
		result.IsLeaf = true
		result.Decision = PluralityValue(parentExamples, goal)
		return result
	}
	if len(examples.ColFreq(goal)) == 1 {
		result.IsLeaf = true
		result.Decision = examples.Col(goal)[0]
		return result
	}
	if len(attributes) == 0 {
		result.IsLeaf = true
		result.Decision = PluralityValue(examples, goal)
		return result
	}
	attr := MostImportant(examples, attributes, goal)
	result.Attribute = attr
	result.Children = make(map[int]DecisionTree)
	attributes = RemoveAttribute(attributes, attr)
	for v, p := range examples.PartitionBy(attr) {
		result.Children[v] = DecisionTreeLearn(p, attributes, goal, examples)
	}
	return result
}

func RemoveAttribute(attributes []string, attr string) []string {
	result := make([]string, 0, len(attributes)-1)
	for _, a := range attributes {
		if a == attr {
			continue
		}
		result = append(result, a)
	}
	return result
}

// MostImportant selects the attribute with the greatest information gain
func MostImportant(examples *dataset.Dataset, attributes []string, goal string) string {
	bestAttr := attributes[0]
	bestGain := Gain(examples, bestAttr, goal)
	for _, attr := range attributes[1:] {
		gain := Gain(examples, attr, goal)
		if gain > bestGain {
			bestGain = gain
			bestAttr = attr
		}
	}
	return bestAttr
}

// PluralityValue selects the most common outcome from a set of examples, breaking ties randomly
func PluralityValue(examples *dataset.Dataset, goal string) int {
	xs := examples.ColFreq(goal)
	bestX, bestN := -1, -1
	for x, n := range xs {
		if n > bestN {
			bestX, bestN = x, n
		}
	}
	return bestX
}

package events

import "golang.org/x/exp/constraints"

func NormalizeCategory[T constraints.Integer](category []T) []Category {
	res := make([]Category, len(category))
	for i, v := range category {
		res[i] = Category(v)
	}
	return res
}

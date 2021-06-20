package backtracking

import (
	"testing"
)

func TestKnapsack(t *testing.T) {
	items := []uint{2, 2, 4, 6, 3}
	n := uint(len(items))
	k := &Knapsack{
		maxW:   0,
		cw:     0,
		w:      10,
		weight: make([]uint, n),
	}
	k.add(0, n, items)
}

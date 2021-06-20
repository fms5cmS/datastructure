package stairs2

import "testing"

func TestStairs(t *testing.T) {
	mem := make(map[int]int)
	steps := Stairs(10, mem)
	t.Log(steps)
}

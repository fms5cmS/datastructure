package stack

import "testing"

func TestIsValid(t *testing.T) {
	s := "[](){}"
	t.Log(isValid(s))
	s = "[()]{}"
	t.Log(isValid(s))
	s = "("
	t.Log(isValid(s))
	s = "(][]()"
	t.Log(isValid(s))
	s = "]"
	t.Log(isValid(s))
}

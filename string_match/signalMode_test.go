package string_match

import "testing"

func TestBF(t *testing.T) {
	tests := []struct {
		main    string
		pattern string
	}{
		{"abcdefghijk", "fgh"},
		{"moasihfnreusdad", "fnre"},
		{"sfhebuirg", "sddsjkg"},
	}
	for _, testcase := range tests {
		t.Log(BF(testcase.main, testcase.pattern))
	}
}

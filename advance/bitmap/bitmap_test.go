package bitmap

import "testing"

func TestBitma(t *testing.T) {
	bitmap := NewBitmap(19)
	t.Logf("%+v", bitmap)
	bitmap.SetBit(13, 1)
	t.Logf("%b", bitmap.data)
	t.Log(bitmap.GetBit(1))
}

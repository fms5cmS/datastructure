package bitmap

type Bitmap struct {
	data    []byte // 保存实际的 bit 数据
	bitsize uint   // 指示该 Bitmap 的 bit 容量
}

func NewBitmap(size uint) *Bitmap {
	if size == 0 {
		size = 0x01 << 32
	} else if remainder := size % 8; remainder != 0 {
		size += 8 - remainder
	}
	return &Bitmap{
		data:    make([]byte, size>>3),
		bitsize: size,
	}
}

// 将 offset 位置的 bit 置为 value (非 0 即判定为 1)
func (bitmap *Bitmap) SetBit(offset uint, value uint) bool {
	if bitmap.bitsize < offset {
		return false
	}
	index, pos := offset/8, offset%8
	if value == 0 {
		// &^ 是 Go 中的按位置零
		bitmap.data[index] &^= 0x01 << pos
	} else {
		bitmap.data[index] |= 0x01 << pos
	}
	return true
}

// 判断 offset 位置的值是否为 1
func (bitmap *Bitmap) GetBit(offset uint) bool {
	index, pos := offset/8, offset%8
	if bitmap.bitsize < offset {
		return false
	}
	return (bitmap.data[index] & (0x01 << pos)) != 0
}

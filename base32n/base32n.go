package base32n

import (
	"errors"
	"strconv"
	"unsafe"
)

var StdEncoding = NewEncoding("234567abcdefghijklmnopqrstuvwxyz")

const invalidIndex = '\xff'

type Encoding struct {
	encode    [32]byte   // mapping of symbol index to symbol byte value
	decodeMap [256]uint8 // mapping of symbol byte value to symbol index
}

func NewEncoding(encoder string) *Encoding {
	e := new(Encoding)
	copy(e.encode[:], encoder)
	for i := range len(e.decodeMap) {
		e.decodeMap[i] = invalidIndex
	}
	for i := range len(encoder) {
		char := encoder[i]
		e.decodeMap[char] = uint8(i)
	}
	return e
}

func (enc *Encoding) Encode(num int64) string {
	result := make([]byte, 0, 13)
	for num > 0 {
		remainder := num & 0x1f
		result = append(result, enc.encode[remainder])
		num >>= 5
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return unsafe.String(unsafe.SliceData(result), len(result))
}

func (enc *Encoding) Decode(s string) (int64, error) {
	var num int64 = 0

	for i := range len(s) {
		char := s[i]
		val := enc.decodeMap[char]
		if val == invalidIndex {
			return 0, errors.New("illegal base32n data at input byte " + strconv.FormatInt(int64(char), 10))
		}
		num += int64(val) * int64(1<<uint(5*(len(s)-1-i)))
	}
	return num, nil
}

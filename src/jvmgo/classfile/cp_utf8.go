package classfile

import (
	"fmt"
	"unicode/utf16"
)

type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

// mutf8 -> utf16 -> utf32 -> string
func decodeMUTF8(bytesArr []byte) string {
	utfLen := len(bytesArr)
	charArr := make([]uint16, utfLen)

	var c, char2, char3 uint16
	count := 0
	chararrCount := 0

	for count < utfLen {
		c = uint16(bytesArr[count])
		if c > 127 {
			break
		}
		count++
		charArr[chararrCount] = c
		chararrCount++
	}

	for count < utfLen {
		c = uint16(bytesArr[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxx xxxx */
			count++
			charArr[chararrCount] = c
			chararrCount++
		case 12, 13:
			/* 110x xxxx  10xx xxxx */
			count += 2
			if count > utfLen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytesArr[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			charArr[chararrCount] = c&0x1F<<6 | char2&0x3F
			chararrCount++
		case 14:
			/* 1110 xxxx  10xx xxxx  10xx xxxx */
			count += 3
			if count > utfLen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytesArr[count-2])
			char3 = uint16(bytesArr[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count-1))
			}
			charArr[chararrCount] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			chararrCount++
		default:
			/* 10xx xxxx,  1111 xxxx */
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}
	// The number of chars produced may be less than utfLen
	charArr = charArr[0:chararrCount]
	runes := utf16.Decode(charArr)
	return string(runes)
}

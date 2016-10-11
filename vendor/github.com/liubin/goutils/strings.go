package goutils

import (
	"fmt"
	"strings"
)

type MaskStringParam struct {
	DisplayCharsCount int
	MaskChar          string
	MaskAtTail        bool
}

// Padding string with given pad
func PadString(src string, length int, pad string) string {
	l := len(src)
	if l >= length {
		return src[0:length]
	}

	ss := make([]string, length-l)

	for i := 0; i < length-l; i++ {
		ss[i] = pad
	}
	return fmt.Sprintf("%s%s", src, strings.Join(ss, ""))
}

// Mask string with given mask char for displaying security info.
func MaskString(str string, param MaskStringParam) string {

	length := len(str)

	if length < param.DisplayCharsCount {
		return str
	}

	if param.MaskAtTail {
		return fmt.Sprintf("%s%s", PadString("", length-param.DisplayCharsCount, param.MaskChar), str[(length-param.DisplayCharsCount):])
	} else {
		return PadString(str[0:param.DisplayCharsCount], length, param.MaskChar)
	}
}

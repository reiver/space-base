package spacebeacon

import (
	"strings"
	"unsafe"
)

// Magic is what is put at the beginning of a UDP message.
//
// It can use used to detect whether a UDP message is "speaking" the SPACE-BEACON protocol or not.
const Magic = "SPACE/0.1\n"

func HasMagic(p []byte) bool {
	var str string = unsafe.String(unsafe.SliceData(p), len(p))
	return strings.HasPrefix(str, Magic)
}

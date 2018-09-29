package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	inputs := [][]byte{
		[]byte{0x01},
		[]byte{0x02},
		[]byte{0x7f},
		[]byte{0x80, 0x01},
		[]byte{0xff, 0x01},
		[]byte{0x80, 0x02},
	}
	for _, b := range inputs {
		x, n := binary.Uvarint(b)
		if n != len(b) {
			fmt.Println("Uvarint did not consume all of in")
		}
		fmt.Println(x)
	}
}

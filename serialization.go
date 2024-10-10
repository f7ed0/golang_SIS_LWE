package sis

import (
	"bytes"
	"encoding/binary"
)

func SerializeInts(ints []int) (buffer []byte) {
	buffer = make([]byte, len(ints)*2)
	interbuff := make([]byte, len(ints)*2)
	compteur := 0
	for _, el := range ints {
		n := binary.PutVarint(interbuff, int64(el))
		for j := range n {
			buffer[compteur] = interbuff[j]
			compteur++
		}
	}
	return
}

func DeserializeInts(buffer []byte, length int) (ints []int, err error) {
	ints = make([]int, length)
	r := bytes.NewBuffer(buffer)
	var j int64
	for i := range ints {
		j, err = binary.ReadVarint(r)
		if err != nil {
			return
		}
		ints[i] = int(j)
	}
	return
}

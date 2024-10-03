package sis

import (
	"bytes"
	"encoding/binary"
)

func SerializeInts(ints []int) (buffer []byte) {
	buffer = make([]byte, len(ints)*4)
	for _, el := range ints {
		binary.PutVarint(buffer, int64(el))
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

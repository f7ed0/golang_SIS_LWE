package sis

import (
	"crypto/rand"
	"fmt"
	_rand "math/rand"
	"testing"
)

func TestCheckShouldNotError(t *testing.T) {
	a, b, err := Default.GenerateCheck([]byte("Test"))
	fmt.Println(a)
	fmt.Println(b)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestCheckAndVerify(t *testing.T) {
	for range 10 {
		var message []byte = make([]byte, _rand.Intn(1024))
		rand.Read(message)
		A_buff, v_buff, err := Default.GenerateCheck(message)
		if err != nil {
			t.Errorf("Error during GenerateCheck : %s", err.Error())
		}
		ok, err := Default.Validate(message, A_buff, v_buff)
		if err != nil {
			t.Errorf("Error during validate : %s\n", err.Error())
		}
		if !ok {
			t.Error("Validate failed.")
		}
	}
}

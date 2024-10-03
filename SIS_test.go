package sis

import (
	"bytes"
	"crypto/rand"
	_rand "math/rand"
	"slices"
	"testing"
)

func TestCheckShouldNotError(t *testing.T) {
	_, _, err := Default.GenerateCheck([]byte("Test"))
	// fmt.Println(a)
	// fmt.Println(b)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestCheckAndVerify(t *testing.T) {
	for range 20 {
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

func TestShouldNotValidate(t *testing.T) {
	for range 20 {
		var message []byte = make([]byte, _rand.Intn(1024))
		rand.Read(message)
		var message2 []byte = make([]byte, _rand.Intn(1024))
		rand.Read(message2)
		for bytes.Equal(message, message2) {
			var message2 []byte = make([]byte, _rand.Intn(1024))
			rand.Read(message2)
		}
		A_buff, v_buff, err := Default.GenerateCheck(message)
		if err != nil {
			t.Errorf("Error during GenerateCheck : %s", err.Error())
		}
		ok, err := Default.Validate(message2, A_buff, v_buff)
		if err != nil {
			t.Errorf("Error during validate : %s\n", err.Error())
		}
		if ok {
			t.Error("Validate sucess but should have failed.")
		}
	}
}

func TestSerializationDeserialization(t *testing.T) {
	for range 20 {
		var message []byte = make([]byte, _rand.Intn(1024))
		rand.Read(message)
		A_buff, v_buff, err := Default.GenerateCheck(message)
		if err != nil {
			t.Errorf("Error during GenerateCheck : %s", err.Error())
		}
		A_bytebuff := SerializeInts(A_buff)
		v_bytebuff := SerializeInts(v_buff)
		A_debuff, err := DeserializeInts(A_bytebuff, Default.N*Default.M)
		if err != nil {
			t.Errorf("Error during deserialization : %s", err.Error())
		}
		v_debuff, err := DeserializeInts(v_bytebuff, Default.N)
		if err != nil {
			t.Errorf("Error during deserialization : %s", err.Error())
		}

		slices.Equal(A_buff, A_debuff)
		slices.Equal(v_buff, v_debuff)
	}
}

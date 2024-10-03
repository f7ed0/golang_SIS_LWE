package matrix

import (
	"fmt"
	"slices"
	"testing"
)

func TestMul(t *testing.T) {
	a := NewZeroMatrix(2, 2)
	a.Set(0, 0, 1)
	a.Set(0, 1, 2)
	a.Set(1, 0, 3)
	a.Set(1, 1, 4)

	b := NewZeroMatrix(2, 1)
	b.Set(0, 0, 5)
	b.Set(1, 0, 6)

	fmt.Println(a, b)

	c, err := a.MulMod(b, 1000000)
	if err != nil {
		t.Error(err.Error())
	}

	byt, err := c.LineToBytes()
	if err != nil {
		t.Error(err.Error())
	}

	if !slices.Equal(byt, []byte{17, 39}) {
		t.Errorf("slice not equal")
	}
}

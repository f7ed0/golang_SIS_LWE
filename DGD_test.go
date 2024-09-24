package sis

import (
	"fmt"
	"testing"
)

func TestGaussianSumShouldBeOne(t *testing.T) {
	d := NewDGD(100, 0.5)
	fmt.Println(d)
	sum := float64(0)
	for _, item := range d {
		sum += item
	}
	fmt.Println(sum)
	if sum < 0.99 && sum > 1.01 {
		t.Errorf("sum should be near 1 (<0.99) an is %v", sum)
	}
}

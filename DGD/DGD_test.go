package dgd

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestRepartition(t *testing.T) {
	const iteration_count int = 100000
	for j := 0; j < 10; j++ {
		q := uint64(rand.Int63n(10000000))
		s := rand.Float64() * 100
		d := NewDGD(q, s)
		var mean float64 = 0
		var stdev float64 = 0
		for i := 0; i < iteration_count; i++ {
			result := d.Rand()
			mean += float64(result) / float64(iteration_count)
			stdev += math.Pow(float64(result), 2) / float64(iteration_count)
		}
		stdev = math.Sqrt(stdev - mean*mean)

		if math.Abs(stdev-s) > 0.02*s || math.Abs(mean-float64(q/2)) > 0.02*float64(q/2) {
			t.Errorf("FAILURE => have (%v %v), wants (%v %v)\n", mean, stdev, float64(q/2), s)
			return
		}
		fmt.Printf("SUCCESS => have (%v %v), wants (%v %v)\n", mean, stdev, float64(q/2), s)
	}
}

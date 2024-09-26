package sis

import (
	"math"
	"time"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

type DiscreteGaussianDistribution struct {
	normal distuv.Normal
	Q      uint64
}

func NewDGD(q uint64, sigma float64) (d DiscreteGaussianDistribution) {
	d = DiscreteGaussianDistribution{
		Q: q,
		normal: distuv.Normal{
			Mu:    float64(q / 2),
			Sigma: sigma,
			Src:   rand.NewSource(uint64(time.Now().UnixNano())),
		},
	}
	return
}

func (d DiscreteGaussianDistribution) Rand() uint64 {
	return uint64(math.Floor(d.normal.Rand())) % d.Q
}

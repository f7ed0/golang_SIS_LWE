package sis

import (
	"time"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

type DiscreteGaussianDistribution struct {
	distuv.Normal
}

func NewDGD(q uint64, sigma float64) (d DiscreteGaussianDistribution) {
	d = DiscreteGaussianDistribution{distuv.Normal{
		Mu:    float64(q / 2),
		Sigma: sigma,
		Src:   rand.NewSource(uint64(time.Now().UnixMilli())),
	}}
	return
}

func (d DiscreteGaussianDistribution) Rand() uint64 {
	v := rand.Float64()
	i := 0.0
	for v >= d.CDF(i) {
		i++
	}
	return uint64(i - 1)
}

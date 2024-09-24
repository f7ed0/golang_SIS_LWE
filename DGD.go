package sis

import (
	"time"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

type DiscreteGaussianDistribution []float64

func NewDGD(q uint64, sigma float64) (d DiscreteGaussianDistribution) {
	d = make(DiscreteGaussianDistribution, q)
	s := rand.NewSource(uint64(time.Now().UnixMilli()))
	normal := distuv.Normal{
		Mu:    float64(q / 2),
		Sigma: sigma,
		Src:   s,
	}
	for i := 0; i < int(q); i++ {
		d[i] = normal.CDF(float64(i+1)) - normal.CDF(float64(i))
	}
	return
}

func (d DiscreteGaussianDistribution) Rand() uint64 {
	v := rand.Float64()
	i := 0
	for v > 0 {
		v -= d[i]
		i++
	}
	return uint64(i)
}

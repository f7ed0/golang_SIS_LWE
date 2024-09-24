package sis

type SasLwe struct {
	N   uint64
	Q   uint64
	Khi DiscreteGaussianDistribution
}

func NewSasLwe(n uint64, q uint64) (result SasLwe) {
	result = SasLwe{
		N:   n,
		Q:   q,
		Khi: NewDGD(q, 0.5),
	}
	return
}

func (s *SasLwe) Encrypt() {

}

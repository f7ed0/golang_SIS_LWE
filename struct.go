package sis

import (
	"crypto/sha512"
	"hash"
	"slices"
	"sync"
	"time"

	"github.com/f7ed0/golang_SIS_LWE/matrix"
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

type SIS struct {
	src    distuv.Uniform
	hasher hash.Hash
	sync.RWMutex
	N int
	Q int
	M int
}

var Default SIS = NewSISSHA512(4099, 64, 1537)

func NewSISSHA512(q, n, m int) (result SIS) {
	result = SIS{
		N:      n,
		Q:      q,
		M:      m,
		hasher: sha512.New(),
		src: distuv.Uniform{
			Src: rand.NewSource(uint64(time.Now().UnixMilli())),
			Min: 0,
			Max: float64(q),
		},
	}
	return
}

func (s *SIS) GenerateRandomInt() int {
	return int(s.src.Rand())
}

func (s *SIS) generateAMaxtrix() matrix.Matrix {
	s.Lock()
	defer s.Unlock()
	A := matrix.NewZeroMatrix(s.N, s.M)
	for i := range s.N {
		for j := range s.M {
			A.Set(i, j, s.GenerateRandomInt())
		}
	}
	return A
}

func (s *SIS) GenerateCheck(m []byte) (A_buff []int, v_buff []int, err error) {
	hashed_m := s.generateHash(m)
	x := matrix.BytesToColumn(hashed_m, s.Q)
	A := s.generateAMaxtrix()
	v, err := A.MulMod(x, s.Q)
	if err != nil {
		return
	}
	v_buff, err = v.LineToInts()
	A_buff = A.MatToInts()
	if err != nil {
		return
	}
	return
}

func (s *SIS) generateHash(m []byte) (hash []byte) {
	s.Lock()
	hash = make([]byte, s.M)
	hashed_m := s.hasher.Sum(m)
	s.hasher.Reset()
	s.Unlock()
	for i := range s.M {
		hash[i] = hashed_m[i%len(hashed_m)]
	}
	return
}

func (s *SIS) Validate(m []byte, A_buff, v_buff []int) (ok bool, err error) {
	ok = false
	hashed_m := s.generateHash(m)
	x := matrix.BytesToColumn(hashed_m, s.Q)
	A, err := matrix.IntsToA(A_buff, s.N, s.M)
	if err != nil {
		return
	}
	v, err := A.MulMod(x, s.Q)
	if err != nil {
		return
	}
	vp_buff, err := v.LineToInts()
	if err != nil {
		return
	}
	ok = slices.Equal(v_buff, vp_buff)
	return
}

package sis

import (
	"crypto/sha512"
	"hash"
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
	n int
	q int
	m int
}

var Default SIS = NewSISSHA512(4099, 64, 1537)

func NewSISSHA512(q, n, m int) (result SIS) {
	result = SIS{
		n:      n,
		q:      q,
		m:      m,
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
	A := matrix.NewZeroMatrix(s.n, s.m)
	for i := range s.n {
		for j := range s.m {
			A.Set(i, j, s.GenerateRandomInt())
		}
	}
	return A
}

func (s *SIS) GenerateCheck(m []byte) (A_buff []int, v_buff []int, err error) {
	hashed_m := s.generateHash(m)
	x := matrix.BytesToColumn(hashed_m, s.q)
	A := s.generateAMaxtrix()
	v, err := A.MulMod(x, s.q)
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
	hash = make([]byte, s.m)
	s.hasher.Reset()
	s.Unlock()
	hashed_m := s.hasher.Sum(m)
	for i := range s.m {
		hash[i] = hashed_m[i%len(hashed_m)]
	}
	return
}

func (s *SIS) Validate(m []byte, A_buff, v_buff []int) (ok bool, err error) {
	ok = false
	hashed_m := s.generateHash(m)
	x := matrix.BytesToColumn(hashed_m, s.q)
	A, err := matrix.IntsToA(A_buff, s.n, s.m)
	if err != nil {
		return
	}
	v, err := A.MulMod(x, s.q)
	if err != nil {
		return
	}
	vp_buff, err := v.LineToInts()
	if err != nil {
		return
	}
	ok = true
	if len(v_buff) != len(vp_buff) {
		return false, nil
	}
	for i := 0; i < len(v_buff) && i < len(vp_buff); i++ {
		if v_buff[i] != vp_buff[i] {
			ok = false
			return
		}
	}
	return
}

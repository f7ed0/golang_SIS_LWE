package matrix

import "errors"

type Matrix struct {
	arr [][]int
	n   int
	m   int
}

func NewZeroMatrix(n, m int) (mat Matrix) {
	mat.arr = make([][]int, n)
	mat.n = n
	mat.m = m
	for i := range mat.arr {
		mat.arr[i] = make([]int, m)
	}
	return
}

func (mat *Matrix) Set(i, j, value int) error {
	if i >= mat.n || j >= mat.m {
		return errors.New("index to big for assignment")
	}
	mat.arr[i][j] = value
	return nil
}

func (mat Matrix) Get(i, j int) (int, error) {
	if i >= mat.n || j >= mat.m {
		return 0, errors.New("index to bin for read")
	}
	return mat.arr[i][j], nil
}

func (mat Matrix) MulMod(o Matrix, modulo int) (result Matrix, err error) {
	if mat.m != o.n {
		println(mat.n, mat.m)
		println(o.n, o.m)
		err = errors.New("can't multiply these matrix")
		return
	}
	result.n = o.m
	result.m = mat.n
	result.arr = make([][]int, result.n)
	for i := range result.arr {
		result.arr[i] = make([]int, result.m)
		for j := range result.arr[i] {
			result.arr[i][j] = 0
			for k := range mat.m {
				result.arr[i][j] = (result.arr[i][j] + mat.arr[j][k]*o.arr[k][i]) % modulo
			}
		}
	}
	return
}

func (m Matrix) LineToInts() (b []int, err error) {
	if m.n > 1 {
		err = errors.New("can't convert to byte a multiple line matrix")
		return
	}

	b = make([]int, m.m)

	for i := range m.m {
		b[i] = m.arr[0][i]
	}

	return
}

func IntsToColumn(b []int, modulo int) (m Matrix) {
	m.arr = make([][]int, len(b))
	m.n = len(b)
	m.m = 1
	for i := range len(b) {
		m.arr[i] = []int{b[i] % modulo}
	}
	return
}

func BytesToColumn(b []byte, modulo int) (m Matrix) {
	m.arr = make([][]int, len(b))
	m.n = len(b)
	m.m = 1
	for i := range len(b) {
		m.arr[i] = []int{int(b[i]) % modulo}
	}
	return
}

func IntsToA(b []int, n, m int) (mat Matrix, err error) {
	if len(b) < n*m {
		err = errors.New("buffer not long enought for n*m matrix generation")
		return
	}
	mat.n = n
	mat.m = m
	mat.arr = make([][]int, n)
	for i := range n {
		mat.arr[i] = make([]int, m)
		for j := range m {
			mat.arr[i][j] = int(b[i*m+j])
		}
	}
	return
}

func (mat Matrix) MatToInts() (b []int) {
	b = make([]int, mat.n*mat.m)
	for i := range mat.n {
		for j := range mat.m {
			b[i*mat.m+j] = mat.arr[i][j]
		}
	}
	return
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	m := setRandom(4, 5, 9)
	fmt.Println(m)
}

func setRandom(dimX, dimY, n int) [][]uint8 {
	v := make([]uint8, dimX*dimY)
	for i := 0; i < n; i++ {
		v[i] = 1
	}
	shuffle(v)
	m := roll(v, dimX, dimY)
	return m
}

func roll(v []uint8, dimX, dimY int) [][]uint8 {
	m := newMatrix(dimX, dimY)
	for i := 0; i < len(v); i++ {
		m[i/dimY][i%dimY] = v[i]
	}
	return m
}

func newMatrix(dimX, dimY int) [][]uint8 {
	_2d := make([][]uint8, dimX)
	for i := 0; i < dimX; i++ {
		_2d[i] = make([]uint8, dimY)
	}
	return _2d
}

func shuffle(v []uint8) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < len(v); i++ {
		j := rand.Intn(len(v))
		if j > i {
			v[i], v[j] = v[j], v[i]
		}
	}
}

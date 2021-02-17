package main

import "math/rand"

func randByte() byte {
	max := 256                  // 255 + 1 because rand.Intn return int >= 0, < n
	return byte(rand.Intn(max)) // will return num in [0, 256)
}

func pixelSetAt(screen [32]uint64, x int, y int) bool {
	columnFilter := uint64(1) << (63 - uint(x))
	return screen[y]&columnFilter == columnFilter
}

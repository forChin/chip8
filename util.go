package main

import "math/rand"

func randByte() byte {
	max := 256                  // 255 + 1 because rand.Intn return int >= 0, < n
	return byte(rand.Intn(max)) // will return num in [0, 256)
}

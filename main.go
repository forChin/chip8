package main

import (
	"math/rand"
	"runtime"
	"time"
)

const (
	windowW = 640
	windowH = 320

	gameROMPath = "./res/roms/PONG2"
)

type word uint16

func init() {
	runtime.LockOSThread()
	rand.Seed(time.Now().UnixNano())
}

func main() {

}

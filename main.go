package main

import (
	"io/ioutil"
	"math/rand"
	"time"
)

const (
	windowW = 640
	windowH = 320
)

type word uint16

var (
	gameMemory     [0x100]byte // 0xfff ?
	registers      [16]byte
	addressI       word
	gameStack      stack
	programCounter word

	// [3] is RGB
	screenData [64][32][3]byte // [32][64] ?
	keyState   []byte

	delayTimer byte
	soundTimer byte
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func cpuReset() error {
	addressI = 0
	programCounter = 0x200

	gameData, err := ioutil.ReadFile("file")
	if err != nil {
		return err
	}
	copy(gameMemory[200:], gameData)

	return nil
}

func pressedKey() int {
	for i, k := range keyState {
		if k != 0 {
			return i
		}
	}

	return -1
}

func main() {
}

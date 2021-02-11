package main

import (
	"io/ioutil"
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
	screenData     [64][32]byte // [32][64] ?
)

func cpuReset() error {
	addressI = 0
	programCounter = 0x200

	game, err := ioutil.ReadFile("file")
	if err != nil {
		return err
	}
	copy(gameMemory[200:], game)

	return nil
}

func main() {
}

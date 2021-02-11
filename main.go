package main

import (
	"fmt"
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
	screenData     [64][32]byte // [32][64] ?
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

func main() {
	fmt.Println(randByte())
	fmt.Println(randByte())
	fmt.Println(randByte())
	fmt.Println(randByte())
	return
	var opcode word = 0x8737

	registers[0x7] = 0x0091 // x
	registers[0x3] = 0x0097 // y

	opcode8XY7(opcode)

	fmt.Println(registers[0x7] == registers[0x3]-0x0091)
	fmt.Println(registers[0xf] == 1)
}

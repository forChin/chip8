package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowW = 640
	windowH = 320

	gameROMPath = "./res/roms/Tank.ch8"
)

type word uint16

var (
	gameMemory     [0xfff]byte
	registers      [16]byte
	addressI       word
	gameStack      stack
	programCounter word

	// [3] is RGB
	screenData [32]uint64 // [32][64] ?
	keyState   [16]byte

	delayTimer byte
	soundTimer byte
	running    bool
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func cpuReset() error {
	addressI = 0
	programCounter = 0x200

	gameData, err := ioutil.ReadFile(gameROMPath)
	if err != nil {
		return err
	}
	copy(gameMemory[0x200:], gameData)

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
	cpuReset()

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatal(err)
	}
	defer sdl.Quit()

	wind, err := sdl.CreateWindow("Emulator", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, windowW, windowH, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}
	defer wind.Destroy()

	surf, err := wind.GetSurface()
	if err != nil {
		log.Fatal(err)
	}
	surf.FillRect(nil, 0)

	// running = true
	// go func() {
	// 	for {
	// 		event := sdl.PollEvent()
	// 		if event != nil {
	// 			switch event.(type) {
	// 			case *sdl.QuitEvent:
	// 				// fmt.Println("QUIT")
	// 				running = false
	// 				break
	// 			}
	// 		}
	// 	}
	// }()

	// runtime.LockOSThread()
	// for running {
	// 	for y := range screenData {
	// 		for x := range screenData[y] {
	// 			color := uint32(screenData[y][x][0])
	// 			if color > 0 {
	// 				color = 0xffff0000
	// 			} else {
	// 				color = 0x0f0f00f0
	// 			}
	// 			rect := sdl.Rect{int32(x), int32(y), 1, 1}

	// 			surf.FillRect(&rect, color)
	// 		}
	// 	}
	// 	wind.UpdateSurface()
	// 	// time.Sleep(time.Second)
	// 	executeNextOpcode()
	// }
}

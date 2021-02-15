package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
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
	screenData [640][320][3]byte // [32][64] ?
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
	fmt.Print(0x0000f == 0x0f)
	return
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}
	defer sdl.Quit()

	wind, err := sdl.CreateWindow("Emulator", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, windowW, windowH, sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}
	defer wind.Destroy()

	surf, err := wind.GetSurface()
	if err != nil {
		return err
	}
	surf.FillRect(nil, 0)

	running := true
	go func() {
		for {
			event := sdl.PollEvent()
			if event != nil {
				switch event.(type) {
				case *sdl.QuitEvent:
					fmt.Println("QUIT")
					running = false
					break
				}
			}
		}
	}()

	for running {
		for y := range screenData {
			for x := range screenData[y] {
				color := screenData[y][x][0]
				rect := sdl.Rect{int32(x), int32(y), 1, 1}

				surf.FillRect(&rect, uint32(color))
			}
		}
		wind.UpdateSurface()
		time.Sleep(time.Second)
	}

	return nil
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"runtime"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowW = 640
	windowH = 320

	gameROMPath = "./res/roms/PONG2"
)

type word uint16

var (
	keyMap = [16]sdl.Keycode{
		sdl.K_x, sdl.K_1, sdl.K_2, sdl.K_3,
		sdl.K_q, sdl.K_w, sdl.K_e, sdl.K_a,
		sdl.K_s, sdl.K_d, sdl.K_z, sdl.K_c,
		sdl.K_4, sdl.K_r, sdl.K_f, sdl.K_v,
	}

	fontSet = [80]byte{
		0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
		0x20, 0x60, 0x20, 0x20, 0x70, // 1
		0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
		0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
		0x90, 0x90, 0xF0, 0x10, 0x10, // 4
		0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
		0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
		0xF0, 0x10, 0x20, 0x40, 0x40, // 7
		0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
		0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
		0xF0, 0x90, 0xF0, 0x90, 0x90, // A
		0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
		0xF0, 0x80, 0x80, 0x80, 0xF0, // C
		0xE0, 0x90, 0x90, 0x90, 0xE0, // D
		0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
		0xF0, 0x80, 0xF0, 0x80, 0x80, // F
	}
)
var (
	gameMemory     [0xfff]byte
	registers      [16]byte
	addressI       word
	gameStack      stack
	programCounter word

	screenData [32][64]byte
	keyState   [16]bool

	delayTimer byte
	soundTimer byte
	running    bool
)

func init() {
	runtime.LockOSThread()
	rand.Seed(time.Now().UnixNano())
}

func cpuReset() error {
	addressI = 0
	programCounter = 0x200

	gameData, err := ioutil.ReadFile(gameROMPath)
	if err != nil {
		return err
	}
	copy(gameMemory[0x200:], gameData) // check len

	for i, f := range fontSet {
		gameMemory[i] = f
	}

	return nil
}

func pressedKey() int {
	for i, k := range keyState {
		if k {
			return i
		}
	}

	return -1
}

func main() {
	cpuReset()

	go startTimers()
	go startMachine()

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

	running = true
	go func() { // not in own goroutine ?
		for {
			event := sdl.PollEvent()
			if event != nil {
				switch t := event.(type) {
				case *sdl.QuitEvent:
					running = false
					return
				case *sdl.KeyboardEvent:
					updateKey(t.Keysym.Sym, t.State == sdl.PRESSED)
				}
			}
		}
	}()

	const scale = 10
	fps := 60
	frameLen := time.Duration(1000/fps) * time.Millisecond

	for running {
		surf.FillRect(nil, 0)
		for y := range screenData {
			for x := range screenData[y] {
				if screenData[y][x] > 0 {
					xCoord := x * scale
					yCoord := y * scale
					rect := sdl.Rect{int32(xCoord), int32(yCoord), scale, scale}
					surf.FillRect(&rect, 0xffffffff)
				}
			}
		}
		wind.UpdateSurface()
		time.Sleep(frameLen)
	}
}

func updateKey(keyCode sdl.Keycode, state bool) {
	for i, k := range keyMap {
		if keyCode == k {
			keyState[i] = state
		}
	}
}

func startTimers() {
	ticker := time.NewTicker(16667 * time.Microsecond) // 60 Hz

	for range ticker.C {
		if running {
			if delayTimer > 0 {
				delayTimer--
			}
			if soundTimer > 0 {
				soundTimer--
			}

			if soundTimer > 0 {
				fmt.Println("BEEP") // make real sound
			}
		}
	}
}

func startMachine() {
	ticker := time.NewTicker(3 * time.Millisecond) // ~300Hz (from config)

	for range ticker.C {
		next := getNextOpcode()
		fmt.Printf("---: 0x%x\n", next)
		executeOpcode(next)
	}

}

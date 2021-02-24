package main

import (
	"log"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type display struct {
	windowW  int32
	windowH  int32
	pixels   [32][64]byte
	scale    int32
	renderer *sdl.Renderer
}

func newDisplay(width, height int32) *display {
	d := display{windowW: width, windowH: height}
	d.scale = int32(len(d.pixels)) / width

	return &d
}

func (d *display) start() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatal(err)
	}
	defer sdl.Quit()

	wind, err := sdl.CreateWindow(
		"Chip8 Emulator", sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED, d.windowW,
		d.windowH, sdl.WINDOW_SHOWN,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer wind.Destroy()

	d.renderer, err = sdl.CreateRenderer(wind, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatal(err)
	}
	defer d.renderer.Destroy()

	d.renderer.Clear()
	d.renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)

	const fps = 60
	frameLen := time.Duration(1000/fps) * time.Millisecond

	ticker := time.NewTicker(frameLen)
	for range ticker.C {
		d.render()
	}
}

func (d *display) render() {
	d.renderer.SetDrawColor(0, 0, 0, 128)
	d.renderer.FillRect(nil)

	for y := range d.pixels {
		for x := range d.pixels[y] {
			if d.pixels[y][x] > 0 {
				xCoord := int32(x) * d.scale
				yCoord := int32(y) * d.scale
				rect := sdl.Rect{xCoord, yCoord, d.scale, d.scale}
				d.renderer.SetDrawColor(255, 255, 255, 255)
				d.renderer.FillRect(&rect)
			}
		}
	}

	d.renderer.Present()
}

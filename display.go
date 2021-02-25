package main

import (
	"log"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type display struct {
	width    int32
	height   int32
	pixels   [32][64]byte
	renderer *sdl.Renderer
}

func newDisplay(w, h int32) *display {
	return &display{width: w, height: h}
}

func (d *display) start() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatal(err)
	}
	defer sdl.Quit()

	wind, err := sdl.CreateWindow(
		"Chip8", sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED, d.width,
		d.height, sdl.WINDOW_SHOWN,
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
	frameLen := time.Second / time.Duration(fps)

	ticker := time.NewTicker(frameLen)
	for range ticker.C {
		d.render()
	}
}

func (d *display) render() {
	d.renderer.SetDrawColor(0, 0, 0, 128)
	d.renderer.FillRect(nil)

	d.renderer.SetDrawColor(30, 255, 30, 255)

	scale := d.height / int32(len(d.pixels))
	for y := range d.pixels {
		for x := range d.pixels[y] {
			if d.pixels[y][x] > 0 {
				xCoord := int32(x) * scale
				yCoord := int32(y) * scale
				rect := sdl.Rect{
					X: xCoord, Y: yCoord,
					W: scale, H: scale,
				}

				d.renderer.FillRect(&rect)
			}
		}
	}

	d.renderer.Present()
}

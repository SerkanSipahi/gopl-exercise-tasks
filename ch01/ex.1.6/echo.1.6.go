package main

/**
 * @Howto:
 * 1.) go build echo.1.5.go
 * 2.) ./echo.1.6 > my-pic.go
 *
 * @info: you can change the last arg of "SetColorIndex" with "blackIndex, redIndex or blueIndex"
 * for creating the gif with different colors
 */

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var green = color.RGBA{0x00, 0xff, 0x00, 1}
var red   = color.RGBA{0xff, 0x00, 0x00, 1}
var blue  = color.RGBA{0x00, 0x00, 0xff, 1}
var black = color.RGBA{0xff, 0xff, 0xff, 1}

var palette = []color.Color{green, black, red, blue}

const(
	blackIndex = 1
	redIndex   = 2
	blueIndex  = 3
)

func main(){

	lissajous(os.Stdout)

}

func lissajous(out io.Writer){

	const(
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64()
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {

		rect := image.Rect(0,0,2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), redIndex)
		}

		phase +=0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)

}
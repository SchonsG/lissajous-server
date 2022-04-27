// Lissajous generates GIF animations of random lissajous figures
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // First palette color
	blackIndex = 1 // Next palette color
)

func lissajous(out io.Writer, cycles float64) {
	const (
		res     = 0.001 // Angular resolution
		size    = 100   // Image canvas covers from [-size..+size]
		nframes = 64    // Number of animation frames
		delay   = 8     // Time between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // Relative frequency of y oscilator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Differences of phases

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // Note: ignoring erros from codification
}

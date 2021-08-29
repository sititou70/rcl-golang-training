// page 25
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cycles, err := strconv.Atoi(r.URL.Query().Get("cycles"))
		if err != nil {
			lissajous(w, 5)
		} else {
			lissajous(w, cycles)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var palette = []color.Color{
	color.Black,
	color.RGBA{0x94, 0x00, 0xD3, 0xFF},
	color.RGBA{0x4B, 0x00, 0x82, 0xFF},
	color.RGBA{0x00, 0x00, 0xFF, 0xFF},
	color.RGBA{0x00, 0xFF, 0x00, 0xFF},
	color.RGBA{0xFF, 0xFF, 0x00, 0xFF},
	color.RGBA{0xFF, 0x7F, 0x00, 0xFF},
	color.RGBA{0xFF, 0x00, 0x00, 0xFF},
}

const (
	bgIndex  = 0
	colorNum = 7
)

func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001 // angular resolution
		size    = 300   // image canvas covers [-size..+size]
		nframes = 512   // number of animation frames
		delay   = 2     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 1.5 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		tMax := float64(cycles) * 2 * math.Pi
		for t := 0.0; t < tMax; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex := math.Floor((t/tMax)*colorNum) + 1
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(colorIndex))
		}
		phase += math.Pi * 2 / nframes
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

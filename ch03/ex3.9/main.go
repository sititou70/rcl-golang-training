// page 69
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

// http://localhost:8000/?centerX=-1.2498&centerY=0.117&zoom=5000&maxIterations=2000
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var (
			centerX       = 0.0
			centerY       = 0.0
			zoom          = 0.5
			width, height = 1024, 1024
			maxIterations = 100
			brightness    = 3
		)

		qCenterX, err := strconv.ParseFloat(r.URL.Query().Get("centerX"), 0)
		if err == nil {
			centerX = qCenterX
		}
		qCenterY, err := strconv.ParseFloat(r.URL.Query().Get("centerY"), 0)
		if err == nil {
			centerY = qCenterY
		}
		qZoom, err := strconv.ParseFloat(r.URL.Query().Get("zoom"), 0)
		if err == nil {
			zoom = qZoom
		}
		qMaxIterations, err := strconv.Atoi(r.URL.Query().Get("maxIterations"))
		if err == nil {
			maxIterations = qMaxIterations
		}

		var (
			xmin, ymin, xmax, ymax = -1/zoom + centerX, -1/zoom - centerY, +1/zoom + centerX, +1/zoom - centerY
		)
		img := image.NewRGBA(image.Rect(0, 0, width, height))
		for py := 0; py < height; py++ {
			y := float64(py)/float64(height)*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/float64(width)*(xmax-xmin) + xmin
				n := mandelbrotComplex128(x, y, maxIterations)

				if n == -1 {
					img.Set(px, py, color.Black)
				} else {
					img.Set(px, py, colorRump(n*brightness, maxIterations))
				}
			}
		}

		png.Encode(w, img) // NOTE: ignoring errors
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func mandelbrotComplex128(x float64, y float64, maxIterations int) int {
	var z = complex(x, y)
	const contrast = 15

	var v complex128
	for n := 0; n < maxIterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return n
		}
	}
	return -1
}

// color util
func mixUint8(v1 uint8, v2 uint8, weight float64) uint8 {
	return v1 + uint8((float64(v2)-float64(v1))*weight)
}
func mixColor(c1 color.RGBA, c2 color.RGBA, weight float64) color.RGBA {
	return color.RGBA{
		R: mixUint8(c1.R, c2.R, weight),
		G: mixUint8(c1.G, c2.G, weight),
		B: mixUint8(c1.B, c2.B, weight),
		A: mixUint8(c1.A, c2.A, weight),
	}
}

func parseRGBA(str string) color.RGBA {
	var black = color.RGBA{
		R: 0,
		G: 0,
		B: 0,
		A: 0,
	}

	if str[0] != '#' {
		return black
	}

	R, err := strconv.ParseInt(str[1:3], 16, 0)
	G, err := strconv.ParseInt(str[3:5], 16, 0)
	B, err := strconv.ParseInt(str[5:7], 16, 0)
	A, err := strconv.ParseInt(str[7:9], 16, 0)
	if err != nil {
		return black
	}

	return color.RGBA{
		R: uint8(R),
		G: uint8(G),
		B: uint8(B),
		A: uint8(A),
	}
}

func colorRump(value int, max int) color.Color {
	var minColor = parseRGBA("#000000FF")
	var maxColor = parseRGBA("#FFFF00FF")

	if value > max {
		return maxColor
	}
	return mixColor(minColor, maxColor, float64(value)/float64(max))
}

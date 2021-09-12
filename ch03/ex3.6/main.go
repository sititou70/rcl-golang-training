// page 69
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"strconv"
)

func main() {
	const (
		centerX                = -1.25
		centerY                = 0.055
		zoom                   = 40.0
		xmin, ymin, xmax, ymax = -2/zoom + centerX, -2/zoom - centerY, +2/zoom + centerX, +2/zoom - centerY
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		dy := 1 / height * (ymax - ymin) / 2
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			dx := 1 / width * (xmax - xmin) / 2
			c1 := mandelbrot(complex(x, y))
			c2 := mandelbrot(complex(x+dx, y))
			c3 := mandelbrot(complex(x, y+dy))
			c4 := mandelbrot(complex(x+dx, y+dy))
			// Image point (px, py) represents complex value z.
			img.Set(px, py, color.RGBA{
				R: uint8((uint64(c1.R) + uint64(c2.R) + uint64(c3.R) + uint64(c4.R)) / 4),
				G: uint8((uint64(c1.G) + uint64(c2.G) + uint64(c3.G) + uint64(c4.G)) / 4),
				B: uint8((uint64(c1.B) + uint64(c2.B) + uint64(c3.B) + uint64(c4.B)) / 4),
				A: 0xFF,
			})
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.RGBA {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return colorRump(int(n), iterations)
		}
	}
	return colorRump(0, iterations)
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

func colorRump(value int, max int) color.RGBA {
	return mixColor(parseRGBA("#000000FF"), parseRGBA("#FFFF00FF"), float64(value)/float64(max))
}

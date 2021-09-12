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
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

// target function: f(z) =  z^4 - 1
//                 f'(z) = 4z^3
//            root: 1, -1, i, -i
// recurrence formula: z_{n+1} = z_{n} - f(z_{n}) / f'(z_{n})
//                             = z_{n} - (z_{n}^4 - 1) / 4z_{n}^3
//                             = (4z_{n}^4 - z_{n}^4 + 1) / 4z_{n}^3
//                             = (3z_{n}^4 + 1) / 4z_{n}^3
//                             = (3z_{n} + (1 / z_{n}^3)) / 4
func newton(start complex128) color.Color {
	const brightness = 7
	const iterationsLimit = 100
	const limitDistance = 0.001

	z := start
	for iteration := 0; iteration < iterationsLimit; iteration++ {
		z = (3*z + (1 / (z * z * z))) / 4

		if cmplx.Abs(1-z) < limitDistance {
			return colorRump(iteration*brightness, iterationsLimit, "#000000FF", "#FF0000FF")
		}
		if cmplx.Abs(-1-z) < limitDistance {
			return colorRump(iteration*brightness, iterationsLimit, "#000000FF", "#00FF00FF")
		}
		if cmplx.Abs(1i-z) < limitDistance {
			return colorRump(iteration*brightness, iterationsLimit, "#000000FF", "#0000FFFF")
		}
		if cmplx.Abs(-1i-z) < limitDistance {
			return colorRump(iteration*brightness, iterationsLimit, "#000000FF", "#FFFF00FF")
		}
	}
	return color.Black
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

func colorRump(value int, max int, minColorStr string, maxColorStr string) color.Color {
	var minColor = parseRGBA(minColorStr)
	var maxColor = parseRGBA(maxColorStr)

	if value > max {
		return maxColor
	}
	return mixColor(minColor, maxColor, float64(value)/float64(max))
}

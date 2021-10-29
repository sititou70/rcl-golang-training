package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"strconv"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		panic(fmt.Sprintf("usage: %s PARALLEL_NUM", os.Args[0]))
	}

	parallelNum, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("invalid argument: %v", err))
	}

	img := drawMandelbrot(defaultDrawingOption, parallelNum)
	png.Encode(os.Stdout, img)
}

type drawingOption struct {
	centerX       float64
	centerY       float64
	zoom          float64
	width, height int
}

var defaultDrawingOption = drawingOption{
	centerX: -1.25,
	centerY: 0.055,
	zoom:    40.0,
	width:   1024,
	height:  1024,
}

func drawMandelbrot(option drawingOption, parallelNum int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, option.width, option.height))

	var wg sync.WaitGroup
	eachHeight := option.height / parallelNum
	for i := 0; i < parallelNum; i++ {
		wg.Add(1)
		if i != parallelNum-1 {
			go drawMandelbrotWithHeightRange(img, option, eachHeight*i, eachHeight*(i+1)-1, &wg)
		} else {
			go drawMandelbrotWithHeightRange(img, option, eachHeight*i, option.height-1, &wg)
		}
	}
	wg.Wait()

	return img
}
func drawMandelbrotWithHeightRange(img *image.RGBA, option drawingOption, hmin, hmax int, wg *sync.WaitGroup) {
	defer wg.Done()

	var (
		xmin = -2/option.zoom + option.centerX
		ymin = -2/option.zoom - option.centerY
		xmax = +2/option.zoom + option.centerX
		ymax = +2/option.zoom - option.centerY
	)

	for py := hmin; py <= hmax; py++ {
		y := float64(py)/float64(option.height)*(ymax-ymin) + ymin
		dy := 1 / float64(option.height) * (ymax - ymin) / 2
		for px := 0; px < option.width; px++ {
			x := float64(px)/float64(option.width)*(xmax-xmin) + xmin
			dx := 1 / float64(option.width) * (xmax - xmin) / 2
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

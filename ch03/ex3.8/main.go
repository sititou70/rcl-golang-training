// page 69
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/big"
	"math/cmplx"
	"os"
	"runtime"
	"strconv"
	"time"
)

type mandelbrotFunc = func(real float64, imaginary float64) (color.Color, int)

func main() {
	var typeSelector = os.Args[1]
	var mandelbrotFunc mandelbrotFunc
	switch typeSelector {
	case "complex64":
		mandelbrotFunc = mandelbrotComplex64
	case "complex128":
		mandelbrotFunc = mandelbrotComplex128
	case "big.Float":
		mandelbrotFunc = mandelbrotBigFloat
	case "big.Rat":
		mandelbrotFunc = mandelbrotBigRat
	default:
		os.Exit(1)
	}

	const (
		centerX                = -1.2498
		centerY                = 0.117
		zoom                   = 100.0
		xmin, ymin, xmax, ymax = -2/zoom + centerX, -2/zoom - centerY, +2/zoom + centerX, +2/zoom - centerY
		width, height          = 256, 256
	)

	// draw mandelbrot
	var start = time.Now().UnixNano()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			color, _ := mandelbrotFunc(x, y)
			img.Set(px, py, color)
		}
	}
	var end = time.Now().UnixNano()
	png.Encode(os.Stdout, img) // NOTE: ignoring errors

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Fprintf(os.Stderr, "time: %v ms\ttotal alloc: %v Bytes\n", float64(end-start)/1000000, mem.TotalAlloc)
}

// mandelbrot functions
func mandelbrotComplex64(x float64, y float64) (color.Color, int) {
	var z = complex64(complex(x, y))
	const iterations = 100
	const contrast = 15

	var v complex64
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			return colorRump(int(n), iterations), n
		}
	}
	return color.Black, -1
}
func mandelbrotComplex128(x float64, y float64) (color.Color, int) {
	var z = complex(x, y)
	const iterations = 100
	const contrast = 15

	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return colorRump(int(n), iterations), n
		}
	}
	return color.Black, -1
}
func mandelbrotBigFloat(x float64, y float64) (color.Color, int) {
	var z = BigFloatComplex{
		r: *big.NewFloat(x),
		i: *big.NewFloat(y),
	}
	const iterations = 100
	const contrast = 15
	var two = big.NewFloat(2)

	var v BigFloatComplex
	for n := 0; n < iterations; n++ {
		v = bigFloatMul(v, v)
		v = bigFloatAdd(v, z)
		abs := bigFloatAbs(v)
		if two.Cmp(&abs) == -1 {
			return colorRump(int(n), iterations), n
		}
	}
	return color.Black, -1
}
func mandelbrotBigRat(x float64, y float64) (color.Color, int) {
	var z = BigRatComplex{
		r: *new(big.Rat).SetFloat64(x),
		i: *new(big.Rat).SetFloat64(y),
	}
	const iterations = 100
	const contrast = 15

	var v BigRatComplex
	for n := 0; n < iterations; n++ {
		v = bigRatMul(v, v)
		v = bigRatAdd(v, z)
		if bigRatAbs(v) > 2 {
			return colorRump(int(n), iterations), n
		}
	}
	return color.Black, -1
}

// bigfloat complex util
type BigFloatComplex = struct {
	r big.Float
	i big.Float
}

func bigFloatAdd(f1 BigFloatComplex, f2 BigFloatComplex) BigFloatComplex {
	var r, i big.Float
	r.Add(&f1.r, &f2.r)
	i.Add(&f1.i, &f2.i)

	return BigFloatComplex{
		r: r,
		i: i,
	}
}

func bigFloatMul(f1 BigFloatComplex, f2 BigFloatComplex) BigFloatComplex {
	var r, i, tmp big.Float

	r.Mul(&f1.r, &f2.r)
	tmp.Mul(&f1.i, &f2.i)
	r.Sub(&r, &tmp)

	i.Mul(&f1.r, &f2.i)
	tmp.Mul(&f1.i, &f2.r)
	i.Add(&i, &tmp)

	return BigFloatComplex{
		r: r,
		i: i,
	}
}

func bigFloatAbs(f1 BigFloatComplex) big.Float {
	var abs, tmp big.Float

	abs.Mul(&f1.r, &f1.r)
	tmp.Mul(&f1.i, &f1.i)
	abs.Add(&abs, &tmp)
	abs.Sqrt(&abs)

	return abs
}

// rat complex util
type BigRatComplex = struct {
	r big.Rat
	i big.Rat
}

func bigRatAdd(f1 BigRatComplex, f2 BigRatComplex) BigRatComplex {
	var r, i big.Rat
	r.Add(&f1.r, &f2.r)
	i.Add(&f1.i, &f2.i)

	return BigRatComplex{
		r: r,
		i: i,
	}
}

func bigRatMul(f1 BigRatComplex, f2 BigRatComplex) BigRatComplex {
	var r, i, tmp big.Rat

	r.Mul(&f1.r, &f2.r)
	tmp.Mul(&f1.i, &f2.i)
	r.Sub(&r, &tmp)

	i.Mul(&f1.r, &f2.i)
	tmp.Mul(&f1.i, &f2.r)
	i.Add(&i, &tmp)

	return BigRatComplex{
		r: r,
		i: i,
	}
}

func bigRatAbs(f1 BigRatComplex) float64 {
	var abs, tmp big.Rat

	abs.Mul(&f1.r, &f1.r)
	tmp.Mul(&f1.i, &f1.i)
	abs.Add(&abs, &tmp)
	float, _ := abs.Float64()
	return math.Sqrt(float)
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

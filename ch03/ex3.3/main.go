// page 67
package main

import (
	"fmt"
	"image/color"
	"math"
	"os"
	"strconv"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type Function = func(x, y float64) float64

func main() {
	var f Function
	switch os.Args[1] {
	case "eggCase":
		f = eggCase
	case "mogul":
		f = mogul
	case "saddle":
		f = saddle
	default:
		os.Exit(0)
	}

	// get max and min value
	var maxValue = math.Inf(-1)
	var minValue = math.Inf(0)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			_, _, z1 := corner(i+1, j, f)
			_, _, z2 := corner(i, j, f)
			_, _, z3 := corner(i, j+1, f)
			_, _, z4 := corner(i+1, j+1, f)
			var z = (z1 + z2 + z3 + z4) / 4
			if z > maxValue {
				maxValue = z
			}
			if z < minValue {
				minValue = z
			}
		}
	}

	// print svg
	var colorRumpSettings = colorRumpSettings{
		maxValue: maxValue,
		minValue: minValue,
		maxColor: parseRGBA("#FF000000"),
		minColor: parseRGBA("#0000FF00"),
	}
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, z1 := corner(i+1, j, f)
			bx, by, z2 := corner(i, j, f)
			cx, cy, z3 := corner(i, j+1, f)
			dx, dy, z4 := corner(i+1, j+1, f)
			var color = colorRump((z1+z2+z3+z4)/4, colorRumpSettings)

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' stroke='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, stringifyRGB(color))
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int, f Function) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	if math.IsInf(z, 0) {
		z = 0
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func eggCase(x, y float64) float64 {
	const scale = 2
	const mergin = 0.1
	// range
	if x > math.Pi*2+mergin || x < -math.Pi*2-mergin {
		return -0.25
	}
	if y > math.Pi*4 || y < -math.Pi*4 {
		return -0.25
	}

	// frame
	if math.Abs((x/6)-math.Trunc(x/6)) < mergin {
		return 0
	}
	if math.Abs((y/6)-math.Trunc(y/6)) < mergin {
		return 0
	}

	var z1 = (-math.Abs(math.Sin(x/scale))+-math.Abs(math.Sin(y/scale)))*0.3 + 0.3
	var z2 float64
	if z1 > 0 {
		z2 = 0
	} else {
		z2 = z1
	}

	return z2
}

func mogul(x, y float64) float64 {
	var z1 = math.Sin(x/1.5) + math.Sin(y/1.5)
	var z2 float64
	if z1 < 0 {
		z2 = z1 * 0.1
	} else {
		z2 = z1 * 0.2
	}
	return (z2) * 0.2
}

func saddle(x, y float64) float64 {
	const scale = 0.1
	return (math.Pow(x*scale, 2) - math.Pow(y*scale, 2)) * scale
}

// color utils
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

type colorRumpSettings struct {
	maxValue float64
	minValue float64
	maxColor color.RGBA
	minColor color.RGBA
}

func colorRump(value float64, settings colorRumpSettings) color.RGBA {
	if value > settings.maxValue {
		return settings.maxColor
	}
	if value < settings.minValue {
		return settings.minColor
	}

	return mixColor(settings.minColor, settings.maxColor, (value-settings.minValue)/(settings.maxValue-settings.minValue))
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

func stringifyRGB(color color.RGBA) string {
	return fmt.Sprintf("#%02x%02x%02x", color.R, color.G, color.B)
}

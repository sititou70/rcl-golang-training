// page 67
package main

import (
	"fmt"
	"math"
	"os"
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

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int, f Function) (float64, float64) {
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
	return sx, sy
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

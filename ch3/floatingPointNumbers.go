package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
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

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	resp := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)
			status := peakValleyOrNeither(az, bz, cz, dz)
			var color string

			if status == "peak" {
				color = "red"
			} else if status == "valley" {
				color = "blue"
			} else {
				color = "white"
			}

			if checkForInfinity([]float64{az, bz, cz, dz}) == false {
				resp += fmt.Sprintf("<polygon style='fill:%s;' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					color, ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	resp += "</svg>"
	w.Write([]byte(resp))
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func checkForInfinity(nums []float64) bool {
	for _, num := range nums {
		if math.IsInf(num, 0) == true {
			return true
		}
	}

	return false
}

func peakValleyOrNeither(z1, z2, z3, z4 float64) string {
	if z1 > 0 && z2 > 0 && z3 > 0 && z4 > 0 {
		return "peak"
	} else if z1 < 0 && z2 < 0 && z3 < 0 && z4 < 0 {
		return "valley"
	}

	return "neither"
}

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"regexp"
)

const (
	width, height            = 600, 320              // canvas size in pixels
	cells                    = 100                   // number of grid cells
	cellRGBColorUnit float64 = 255 / float64(height) // cell rgb color unit
	xyrange                  = 30.0                  // axis ranges (-xyrange..+xyrange)
	xyscale                  = width / 2 / xyrange   // pixels per x or y unit
	zscale                   = height * 0.4          // pixels per z unit
	angle                    = math.Pi / 6           // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {

	/** 3.4 **/ // server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		/** 3.4 **/ // set header
		w.Header().Set("Content-Type", "image/svg+xml")
		createSvg(w)
	})

	log.Fatal(http.ListenAndServe("localhost:22117", nil))
}

func createSvg(w io.Writer) {

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i <= cells; i++ {

		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			fmt.Println("\n", ax, bx, cx, dx)
			fmt.Println("\n", ay, by, cy, dy)

			red := float64(j) * cellRGBColorUnit
			green := float64(j-100) * cellRGBColorUnit

			polygon := fmt.Sprintf("<polygon style='fill: rgb(%g, 0, %g)' points='%g,%g %g,%g %g,%g %g,%g'/>",
				red, green, ax, ay, bx, by, cx, cy, dx, dy)

			// /*3.1*/ Exercise skip non-finite float64
			matched, err := regexp.MatchString("NaN", polygon)
			if err != nil {
				log.Fatalln(err)
			}
			if matched {
				continue
			}
			// /*3.1*/ done

			fmt.Fprint(w, polygon)
		}
	}
	fmt.Fprint(w, "</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-

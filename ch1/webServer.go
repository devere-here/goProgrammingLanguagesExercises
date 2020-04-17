package main

import (
	"log"
	"net/http"
	"io"
	"math/rand"
	"image"
	"math"
	"image/gif"
	"image/color"
	"fmt"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		stuff := r.URL.Query()
		fmt.Println("stuff", stuff, "stuff.cycles", stuff)

		// cycles := 5
		// if stuff.cycles != nil {
		// 	queryCycle, err := stuff.cycles
		// 	fmt.Println("queryCycle", queryCycle, "err", err)
		// }

		lissajous(w, r)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var palette = []color.Color{color.White, color.Black}
const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func lissajous(out io.Writer, req *http.Request) {
	const (
		cycles = 5 // number of complete x oscillator revolutions
		res = 0.001 // angular resolution
		size = 100 // image canvas covers [-size..+size]
		nframes = 64 // number of animation frames
		delay = 8 // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	fmt.Println("cycles ate", cycles)

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
			blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

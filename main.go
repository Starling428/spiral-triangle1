package main

import (
	"flag"
	"fmt"
	"github.com/StephaneBunel/bresenham"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func main() {
	var x1, y1, x2, y2, x1t, x2t, y1t, y2t int
	var linecolor, bgcolor color.Gray16

	degree := flag.Float64("degree", 63, "Rotate angle (degrees)")
	iter := flag.Int("iter", 55, "Quantity of iterations")
	length := flag.Int("length", 30, "Length of the base line")
	size := flag.Int("size", 300, "Length of the image side (the image is a square")
	scale := flag.Float64("scale", 1.06, "Scale multiplier")
	inversion := flag.Bool("inversion", false, "Invert colors")
	name := flag.String("name", "out.png", "Invert colors")

	flag.Parse()

	y2 = *length
	halfsize := *size / 2
	rad := (*degree) * (math.Pi / 180)

	if *inversion {
		linecolor = color.White
		bgcolor = color.Black
	} else {
		linecolor = color.Black
		bgcolor = color.White
	}

	img := image.NewRGBA(image.Rect(0, 0, *size, *size))
	for y := 1; y <= *size; y++ {
		for x := 1; x <= *size; x++ {
			img.Set(x, y, bgcolor)
		}
	}

	for i := 1; i < *iter; i++ {
		bresenham.Bresenham(img, x1+halfsize, y1+halfsize, x2+halfsize, y2+halfsize, linecolor)

		x1t = x1 - x2
		y1t = y1 - y2

		x2t = int((float64(x1t)*math.Cos(rad)-float64(y1t)*math.Sin(rad))*100) / 100
		y2t = int((float64(x1t)*math.Sin(rad)+float64(y1t)*math.Cos(rad))*100) / 100

		x2t = int(float64(x2t)**scale*100) / 100
		y2t = int(float64(y2t)**scale*100) / 100

		x1 = x2
		y1 = y2
		x2 = x2t + x2
		y2 = y2t + y2
	}

	f, err := os.OpenFile(*name, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Can't create a file")
		os.Exit(10)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("Can't close a file")
			os.Exit(20)
		}
	}(f)
	err = png.Encode(f, img)
	if err != nil {
		fmt.Println("Can't encode an image")
		os.Exit(30)
	}
}

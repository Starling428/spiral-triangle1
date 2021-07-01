package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	// Create an 100 x 50 image
	img := image.NewRGBA(image.Rect(0, 0, 1000, 1000))

/*	for i:=0.0; i < 2*math.Pi*170; i+=0.001 {
		img.Set(int(i),   int(math.Sin(0.035*i)*350+500)	  , color.RGBA{R: 255, A: 255})
		img.Set(int(i),   int(math.Sin(0.035*i)*350+500)+1 , color.RGBA{R: 255, A: 255})
		img.Set(int(i),   int(math.Sin(0.035*i)*350+500)+2 , color.RGBA{R: 255, A: 255})
		img.Set(int(i)+1, int(math.Sin(0.035*i)*350+500)   , color.RGBA{R: 255, A: 255})
		img.Set(int(i)+2, int(math.Sin(0.035*i)*350+500)   , color.RGBA{R: 255, A: 255})
		//fmt.Println(fmt.Sprint(int(math.Sin(i)*40+100)))
	}
*/
	Bresenham_1(img, 10, 10, 50, 50, color.RGBA{B: 255, A: 255})

	// Save to out.png
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}


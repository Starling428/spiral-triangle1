package main

import (
	"fmt"
	//"./bresenham"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"strconv"
)

func main() {

	x1 := -15
	x2 := 15
	//x3 := 10
	y1 := -15
	y2 := 15
	//y3 := -10
	var x1t, x2t, y1t, y2t /*, x3t, y3t*/ int

	height, width := 300, 300
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y <= height; y++ {
		for x := 0; x <= width; x++ {
			img.Set(x, y, color.White)
		}
	}
	fmt.Println(os.Args[0])
	if len(os.Args) < 2 {
		os.Exit(420)
	}

	//col := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	/*	for i:=0.0; i < 2*math.Pi*170; i+=0.001 {
			img.Set(int(i),   int(math.Sin(0.035*i)*350+500)	  , color.RGBA{R: 255, A: 255})
			img.Set(int(i),   int(math.Sin(0.035*i)*350+500)+1 , color.RGBA{R: 255, A: 255})
			img.Set(int(i),   int(math.Sin(0.035*i)*350+500)+2 , color.RGBA{R: 255, A: 255})
			img.Set(int(i)+1, int(math.Sin(0.035*i)*350+500)   , color.RGBA{R: 255, A: 255})
			img.Set(int(i)+2, int(math.Sin(0.035*i)*350+500)   , color.RGBA{R: 255, A: 255})
			//fmt.Println(fmt.Sprint(int(math.Sin(i)*40+100)))
		}
	*/
	degree, _ := strconv.ParseFloat(os.Args[1], 32)
	iter, _ := strconv.ParseInt(os.Args[2], 10, 32)
	scale, _ := strconv.ParseFloat(os.Args[3], 32)
	fmt.Println(os.Args[:])

	rad := degree * (math.Pi / 180)
	for i := 1; i < int(iter); i++ {
		Bresenham(img, x1+width/2, y1+height/2, x2+width/2, y2+height/2, color.Black)
		x1t = x1 - x2
		y1t = y1 - y2
		x2t = int((float64(x1t)*math.Cos(rad)-float64(y1t)*math.Sin(rad))*100) / 100
		y2t = int((float64(x1t)*math.Sin(rad)+float64(y1t)*math.Cos(rad))*100) / 100
		//if i%2==0 {
		x2t = x2t * (7 + i) / (6 + i) * int(scale/100)
		y2t = y2t * (7 + i) / (6 + i) * int(scale/100)
		//}
		x1 = x2
		y1 = y2
		x2 = x2t + x2
		y2 = y2t + y2
	}
	/*	for j := 1; j < 55; j++ {
			Bresenham(img, x1+500, y1+500, x2+500, y2+500, color.Black)
			Bresenham(img, x2+500, y2+500, x3+500, y3+500, color.Black)
			Bresenham(img, x3+500, y3+500, x1+500, y1+500, color.Black)

			Bresenham(img, x1+501, y1+501, x2+501, y2+501, color.Black)
			Bresenham(img, x2+501, y2+501, x3+501, y3+501, color.Black)
			Bresenham(img, x3+501, y3+501, x1+501, y1+501, color.Black)

			Bresenham(img, x1+502, y1+502, x2+502, y2+502, color.Black)
			Bresenham(img, x2+502, y2+502, x3+502, y3+502, color.Black)
			Bresenham(img, x3+502, y3+502, x1+502, y1+502, color.Black)

			x1t = int((float64(x1)*math.Cos(rad)-float64(y1)*math.Sin(rad))*100) / 100
			y1t = int((float64(x1)*math.Sin(rad)+float64(y1)*math.Cos(rad))*100) / 100

			x2t = int((float64(x2)*math.Cos(rad)-float64(y2)*math.Sin(rad))*100) / 100
			y2t = int((float64(x2)*math.Sin(rad)+float64(y2)*math.Cos(rad))*100) / 100

			x3t = int((float64(x3)*math.Cos(rad)-float64(y3)*math.Sin(rad))*100) / 100
			y3t = int((float64(x3)*math.Sin(rad)+float64(y3)*math.Cos(rad))*100) / 100

			fmt.Println(x1, " ", x1t, " ", y1, " ", y1t)
			fmt.Println(x2, " ", x2t, " ", y2, " ", y2t)

			x1 = int(float64(x1t)*(1.25)*100) / 100
			x2 = int(float64(x2t)*(1.25)*100) / 100
			y1 = int(float64(y1t)*(1.25)*100) / 100
			y2 = int(float64(y2t)*(1.25)*100) / 100
			x3 = int(float64(x3t)*(1.25)*100) / 100
			y3 = int(float64(y3t)*(1.25)*100) / 100
		}
	*/
	/*
		for j:=1; j<9; j++ {

			Bresenham(img, x1+500, y1+500, x2+500, y2+500, color.Black)
			//Bresenham(img, x2+500, y2+500, x3+500, y3+500, color.Black)
			//Bresenham(img, x3+500, y3+500, x1+500, y1+500, color.Black)

		Bresenham(img, x1+501, y1+501, x2+501, y2+501, color.Black)
			//Bresenham(img, x2+501, y2+501, x3+501, y3+501, color.Black)
			//Bresenham(img, x3+501, y3+501, x1+501, y1+501, color.Black)

			Bresenham(img, x1+502, y1+502, x2+502, y2+502, color.Black)
			//Bresenham(img, x2+502, y2+502, x3+502, y3+502, color.Black)
			//Bresenham(img, x3+502, y3+502, x1+502, y1+502, color.Black)


			//x1=int(float64(x1t)*(1.25)*100)/100
			//x2=int(float64(x2t)*(1.25)*100)/100
			//y1=int(float64(y1t)*(1.25)*100)/100
			//y2=int(float64(y2t)*(1.25)*100)/100

			if j%2 == 0 {
				x1t = int((float64(x1)*math.Cos(rad) - float64(y1)*math.Sin(rad))*100)/100
				y1t = int((float64(x1)*math.Sin(rad) + float64(y1)*math.Cos(rad))*100)/100

				x1=x1t
				y1=y1t
				if j%3 == 0 {
					x1 = int(float64(x1t/4-x2)*(5)*100) / 100
					y1 = int(float64(y1t/4-y2)*(5)*100) / 100
				}
			} else {
				x2t = int((float64(x1)*math.Cos(rad) - float64(y1)*math.Sin(rad))*100)/100
				y2t = int((float64(x1)*math.Sin(rad) + float64(y1)*math.Cos(rad))*100)/100

				x2=x2t
				y2=y2t
				if j%3 == 0 {
					x2 = int(float64(x2t/4-x2)*(5)*100) / 100
					y2 = int(float64(y2t/4-y2)*(5)*100) / 100
				}
			}
			//x2=x1+x2
			//x1=x2-x1
			//x2=x2-x1
			//y2=y1+y2
			//y1=y2-y1
			//y2=y2-y1
			//x3t = int((float64(x3)*math.Cos(rad) - float64(y3)*math.Sin(rad))*100)/100
			//y3t = int((float64(x3)*math.Sin(rad) + float64(y3)*math.Cos(rad))*100)/100

			fmt.Println(x1, " ", x1t, " ", y1, " ", y1t)
			fmt.Println(x2, " ", x2t, " ", y2, " ", y2t)


			//x3=int(float64(x3t)*(1.25)*100)/100
			//y3=int(float64(y3t)*(1.25)*100)/100
		}
	*/

	/*	for j:=1.0; j<15.0; j+=1.0 {
		Bresenham(img, x1t-500, y1t-500, x2t-500, y2t-500, col)

		x1t = int(float64(x1)*math.Cos(0.3*j) - float64(y1)*math.Sin(0.3*j))
		y1t = int(float64(x1)*math.Sin(0.3*j) + float64(y1)*math.Cos(0.3*j))

		x2t = int(float64(x2)*math.Cos(0.3*j) - float64(y2)*math.Sin(0.3*j))
		y2t = int(float64(x2)*math.Sin(0.3*j) + float64(y2)*math.Cos(0.3*j))
		fmt.Println(x1, " ", x1t, " ", y1, " ", y1t)
		fmt.Println(x2, " ", x2t, " ", y2, " ", y2t)

		//x1t=int(float64(x1t)*(5))
		//x2t=int(float64(x2t)*(5))
		//y1t=int(float64(y1t)*(5))
		//y2t=int(float64(y2t)*(5))

	}*/
	/*a, b	:= 100,100
	a1, b1 	:= a,b
	degree 	:= 5.0
	rad		:= degree*(math.Pi/180)

		img.Set(a, b, col)
	for iter:= 1; iter<50; iter++ {
		a1 = int((float64(a)*math.Cos(rad) - float64(b)*math.Sin(rad))*100)/100
		b1 = int((float64(a)*math.Sin(rad) + float64(b)*math.Cos(rad))*100)/100
		img.Set(a+500, b+500, col)
		a=a1
		b=b1
		fmt.Println(a1," ", b1)
	}*/
	// Save to out.png
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}

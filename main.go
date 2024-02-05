package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"math/rand"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Please specify the path to save the generated image")
		os.Exit(1)
	}
	imgPath := os.Args[1]

	out, err := os.Create(imgPath)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer out.Close()

	width, height := rand.Intn(640), rand.Intn(480)
	background := color.RGBA{uint8(rand.Int31n(256)), uint8(rand.Int31n(256)), uint8(rand.Int31n(256)), uint8(rand.Int31n(256))}

	img := createImage(width, height, background)
	for i := 0; i < width*height; i++ {
		img.SetRGBA(i%width, i/width, color.RGBA{uint8(rand.Int31n(256)), uint8(rand.Int31n(256)), uint8(rand.Int31n(256)), uint8(rand.Int31n(256))})
	}

	if strings.HasSuffix(strings.ToLower(imgPath), ".jpg") {
		var opt jpeg.Options
		opt.Quality = 40
		err = jpeg.Encode(out, img, &opt)
	} else {
		panic("Unsupported image format")
	}

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Break image
	image, err := os.ReadFile(imgPath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	brokenImage := make([]byte, len(image))
	for i := 0; i < len(image); i++ {
		brokenImage[i] = image[len(image)-i-1]
	}

	err = os.WriteFile(imgPath+".broken.jpg", brokenImage, 0644)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Printf("Image saved to %s\n", imgPath)
}

func createImage(width int, height int, background color.RGBA) *image.RGBA {
	rect := image.Rect(0, 0, width, height)
	img := image.NewRGBA(rect)
	draw.Draw(img, img.Bounds(), &image.Uniform{background}, image.ZP, draw.Src)
	return img
}

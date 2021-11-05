package main

import (
	"image/color"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func GetImageData(filename string) [][]color.Color {
	image := loadImage(filename)
	return getPixels(image)
}

func loadImage(filename string) image.Image {

	imageFile, openError := os.Open(filename)

	if openError != nil {
		log.Fatal("loadImage failed: " + openError.Error())
	}

	image, _, decodeError := image.Decode(imageFile)

	if decodeError != nil {
		log.Fatal("loadImage failed" + decodeError.Error())
	}

	return image
}

func getPixels(image image.Image) [][]color.Color {

	bounds := image.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y

	var pixels [][]color.Color
	for y := 0; y < height; y++ {
		var row []color.Color
		for x := 0; x < width; x++ {
			row = append(row, image.At(x, y))
		}

		pixels = append(pixels, row)
	}

	return pixels
}

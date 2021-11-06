package media

import (
	"image"
	"image/color"
	_ "image/gif"
	_gif "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"os"
)

func GetFileFormat(file *os.File) string {
	//Thanks stackoverflow lookalike website
	buffer := make([]byte, 512)

	_, err := file.Read(buffer)
	if err != nil {
		log.Fatal("GetFileFormat: " + err.Error())
	}

	contentType := http.DetectContentType(buffer)
	file.Seek(0, 0)

	return contentType
}

func LoadFile(filename string) *os.File {

	file, openError := os.Open(filename)

	if openError != nil {
		log.Fatal("loadImage failed: " + openError.Error())
	}

	return file
}

func decodeImage(file *os.File) image.Image {
	img, _, decodeError := image.Decode(file)

	if decodeError != nil {
		log.Fatal("DecodeImage failed" + decodeError.Error())
	}

	return img
}

func decodeGIF(file *os.File) _gif.GIF {

	gif, err := _gif.DecodeAll(file)

	if err != nil {
		log.Fatal("DecodeGIF error: " + err.Error())
	}

	return *gif
}

type frame struct {
	Pixels [][]color.Color
}

func getFrameImage(img image.Image) *frame {

	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y

	var pixels [][]color.Color
	for y := 0; y < height; y++ {
		var row []color.Color
		for x := 0; x < width; x++ {
			row = append(row, img.At(x, y))
		}

		pixels = append(pixels, row)
	}

	return &frame{Pixels: pixels}
}

func getFrameGIF(img image.Paletted) *frame {

	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y

	var pixels [][]color.Color
	for y := 0; y < height; y++ {
		var row []color.Color
		for x := 0; x < width; x++ {
			row = append(row, img.At(x, y))
		}

		pixels = append(pixels, row)
	}

	return &frame{Pixels: pixels}
}

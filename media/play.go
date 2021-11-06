package media

import (
	"image"
	_gif "image/gif"
	"log"
	"time"

	"github.com/philipjohanszon/GoASCII/characters"
)

func Play(filename string) {

	file := loadFile(filename)
	format := getFileFormat(file)

	switch format {
	case "image/png", "image/jpeg":
		img := decodeImage(file)
		renderImage(img)
	case "image/gif":
		gif := decodeGIF(file)
		playGIF(gif)
	default:
		log.Fatal("Unsupported file format")
	}
}

func playGIF(gif _gif.GIF) {
	var pixelizedImages []*frame

	for _, img := range gif.Image {
		imgFrame := getFrameGIF(*img)
		pixelizedImages = append(pixelizedImages, imgFrame)
	}

	for {
		for _, frame := range pixelizedImages {
			characters.Print(frame.Pixels)
			//just time between frames, about 15 fps
			time.Sleep(67 * time.Millisecond)
		}
	}

}

func renderImage(img image.Image) {

	frame := getFrameImage(img)
	characters.Print(frame.Pixels)
}

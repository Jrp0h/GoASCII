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
	var printFrames []*characters.PrintFrame

	for _, img := range gif.Image {
		imgFrame := getFrameGIF(*img)
		printFrames = append(printFrames, characters.PreRenderFrame(imgFrame.Pixels))
	}

	for {
		for _, frame := range printFrames {
			characters.Print(frame)
			//just time between frames, about 15 fps
			time.Sleep(67 * time.Millisecond)
		}
	}

}

func renderImage(img image.Image) {

	frame := getFrameImage(img)
	printFrame := characters.PreRenderFrame(frame.Pixels)
	characters.Print(printFrame)
}

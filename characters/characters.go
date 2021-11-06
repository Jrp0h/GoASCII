package characters

import (
	"fmt"
	"image/color"
	"os"
	"os/exec"
	"runtime"
)

var characters = [32]string{"'", "¨", ".", "-", "~", "^", "+", "=", "*", "<", "c", "!", "/", "¤", "#", "x", "e", "d", "?", "2", "L", "E", "{", "€", "£", "X", "&", "§", "%", "M", "$", "@"}

type PrintFrame struct {
	Text string
}

func PreRenderFrame(pixels [][]color.Color) *PrintFrame {
	printFrame := PrintFrame{Text: ""}

	for _, row := range pixels {
		for _, pixel := range row {
			printFrame.Text += getCharacterFromColor(pixel)
		}
		printFrame.Text += "\n"
	}

	return &printFrame

}

func Print(frame *PrintFrame) {
	clearScreen()
	fmt.Print(frame.Text)
}

func getCharacterFromColor(pixel color.Color) string {
	r, g, b, _ := pixel.RGBA()
	averageRGBValue := (r + g + b) / 3

	//divide averageRBGValue so that it corresponds to the index of the character array
	//and return the character at that index (16 bit per channel 0-65534)
	return characters[averageRGBValue/2048]
}

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

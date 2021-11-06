package characters

import (
	"fmt"
	"image/color"
	"os"
	"os/exec"
	"runtime"
)

var characters = [32]string{"'", "¨", ".", "-", "~", "^", "+", "=", "*", "<", "c", "!", "/", "¤", "#", "x", "e", "d", "?", "2", "L", "E", "{", "€", "£", "X", "&", "§", "%", "M", "$", "@"}

func Print(pixels [][]color.Color) {
	printString := ""

	for _, row := range pixels {
		for _, pixel := range row {
			printString += getCharacterFromColor(pixel)
		}
		printString += "\n"
	}

	clearScreen()
	fmt.Print(printString)
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

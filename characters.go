package main

import (
	"fmt"
	"image/color"
)

var characters = [32]string{"'", "¨", ".", "-", "~", "^", "+", "=", "*", "<", "c", "!", "/", "¤", "#", "x", "e", "d", "?", "2", "L", "E", "{", "€", "£", "X", "&", "§", "%", "M", "$", "@"}

func PrintCharacters(pixels [][]color.Color) {
	for _, row := range pixels {
		for _, pixel := range row {
			fmt.Print(getCharacterFromColor(pixel))
		}
		fmt.Print("\n")
	}
}

func getCharacterFromColor(pixel color.Color) string {
	r, g, b, _ := pixel.RGBA()
	averageRGBValue := (r + g + b) / 3

	//divide averageRBGValue so that it corresponds to the index of the character array
	//and return the character at that index (16 bit per channel 0-65534)
	return characters[averageRGBValue/2048]
}


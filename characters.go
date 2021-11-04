package main

import (
	"color"
)

var characters := [10]rune{"#", "<", "*", "=", "+", "^", "~", "-", "¨", "'"}


func PrintCharacters(pixels [][]color.Color) {
    for _, row := range pixels {
        for _, pixel := range row {
            fmt.Println(pixel)
        }
    }
}

func getCharacter(pixel color.Color) rune {
    
}

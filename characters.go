package main

import (
    "image/color"
"fmt"
)

var characters = [10]string{"#", "<", "*", "=", "+", "^", "~", "-", "¨", "'"}


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
    //and return the character at that index
    return characters[averageRGBValue/6554]
}
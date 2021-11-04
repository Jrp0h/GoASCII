package main

import (
	"os"
)

func main() {
	//Had to make it one line, just because I could :D
	PrintCharacters(GetImageData(os.Args[1]))
}

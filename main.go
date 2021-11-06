package main

import (
	"github.com/philipjohanszon/GoASCII/media"
	"log"
	"os"
)

func main() {
	//Had to make it one line, just because I could :D
	filename := os.Args[1]

	if len(filename) == 0 {
		log.Fatal("No filename was specified")
	}

	media.Play(filename)
}

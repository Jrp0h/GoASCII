package main

import (
	"github.com/philipjohanszon/GoASCII/media"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No filename was specified")
	}

	//Had to make it one line, just because I could :D
	filename := os.Args[1]
	media.Play(filename)
}

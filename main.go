package main

import (
	"log"

	"github.com/BrianGreenhill/note/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

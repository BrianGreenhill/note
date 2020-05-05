package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const TSFORMAT="20060102150405"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalf("You need to provide a note")
	}

    timestamp := time.Now()
    note := []byte(args[0])
	filename := fmt.Sprintf("/tmp/%s_note", timestamp.Format(TSFORMAT))
	err := ioutil.WriteFile(filename, note, 0644)
	if err != nil {
		log.Fatalf("Could not write file %+v", err)
	}

	fmt.Printf("Taking your note: %s. Find it here: %s", args[0], filename)
}

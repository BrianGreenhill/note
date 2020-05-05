package main

import (
	"fmt"
	"io/ioutil"

    "github.com/briangreenhill/note/config"
)

func main() {
    // get the notes directory from the config
    notesDir := config.Config.NotesDirectory()

    // list the notes in the directory sorted descending by date
    files, err := ioutil.ReadDir(notesDir)
    if err != nil {
        fmt.Printf("Could not read from notes directory %+v", err)
    }

    fmt.Println("Your Notes:")
    for i, f := range files {
        fmt.Printf("[%d] %s\n", i, f.Name())
    }

    // nice: allow sorting by other fields than date
}

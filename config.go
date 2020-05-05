package main

import (
    "fmt"
    "os"
)

type Config struct {
    notesDirectory string
}

const notesDirName = "notes"

// NotesDirectory returns the directory where notes are stored
func (c* Config) NotesDirectory() (string) {
    c.notesDirectory = fmt.Sprintf("%s/.%s", os.Getenv("HOME"), notesDirName)
    return c.notesDirectory
}

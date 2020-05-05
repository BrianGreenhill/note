# notes

To keep track of all those notes you keep taking and losing

## Requirements
Golang

## TODO
- list notes
- add note
- help command
- add repo command
- configuration via config file
- append to existing note
- finish this readme

## How to Build

## Usage

### Directories

notes will make a directory in the home folder given by the `XDG_CONFIG_HOME`
variable for configuration. `NOTES_DIR` tells notes where to put the notes
repository.

### Basic Usage

```
notes {cmd} {note}

notes help = help
    Prints the help menu

notes ls/list
    Lists notes

notes add {note}
    Stores new note in file

notes append {index} {note}
    Appends note to the note given by the index

notes delete {index}
    Deletes the note

notes sync
    Commits new notes to repo
    git add . && git commit && git push origin master
```

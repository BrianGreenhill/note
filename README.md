# note

To keep track of all those notes you keep taking and losing

## Requirements

- Golang >= 1.14
- Notes to take

## TODO

- append to existing note
- search through notes by text
- finish this readme

## Usage

`make` - builds to `bin/note`
`make install` - copies `bin/note` to `$HOME/.local/bin`
`make clean` - truncates the `bin` directory

### Directories

Config is in $HOME/.note.yaml unless specified otherwise by passing the
`cfgFile` flag to the `note` command. Notes are saved in the directory
specified in your config file as `notedir`.

### Basic Usage

```
note {cmd} {note}

note list
    Lists notes

note take {note}
    Stores new note in file

note edit -s {search}
    Searches notes for a given string and edits selected note

note diff
    Displays any notes that need to be synced

note sync
    Commits new note to repo
    git add . && git commit && git push origin master
```

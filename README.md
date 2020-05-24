# note

To keep track of all those notes you keep taking and losing

## Requirements
- Golang >= 1.14
- Notes to take

## TODO
- append to existing note
- define repository
- finish this readme

## How to Build

## Usage

### Directories

Config is in $HOME/.note.yaml unless specified otherwise by passing the `cfgFile` flag to the `note` command. Notes are saved in the directory specified in your config file as `notedir`.

### Basic Usage

```
note {cmd} {note}

note list
    Lists notes

note take {note}
    Stores new note in file

note sync
    Commits new note to repo
    git add . && git commit && git push origin master
```

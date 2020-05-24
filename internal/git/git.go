package git

import (
	"os"
	"os/exec"
	"time"

	"github.com/spf13/viper"
)

func Add() bool {
	return run(exec.Command("git", "add", "."))
}

func Commit() bool {
	return run(exec.Command("git", "commit", "-m", "Updated at: "+time.Now().String()))
}

func Push() bool {
	return run(exec.Command("git", "push"))
}

func Pull() bool {
	return run(exec.Command("git", "pull"))
}

func StashPush() bool {
	return run(exec.Command("git", "stash"))
}

func StashPop() bool {
	return run(exec.Command("git", "stash", "pop"))
}

func run(c *exec.Cmd) bool {
	c.Dir = viper.GetString("notedir")

	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Run()

	return c.ProcessState.Success()
}

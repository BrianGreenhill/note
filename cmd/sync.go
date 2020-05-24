package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/BrianGreenhill/note/internal/git"
)

func init() {
	rootCmd.AddCommand(syncCmd)
}

var (
	syncCmd = &cobra.Command{
		Use:   "sync",
		Short: "Sync your notes",
		Run: func(cmd *cobra.Command, args []string) {
			if !(git.StashPush() && git.Pull() && git.StashPop()) {
				fmt.Println("Could not pull down repo")
			}
			if !(git.Add() && git.Commit()) {
				fmt.Println("Could not commit changes")
			}
			if !(git.Push()) {
				fmt.Println("Could not push changes to repo")
			}
			os.Exit(1)
		},
	}
)

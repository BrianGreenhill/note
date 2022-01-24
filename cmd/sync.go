package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/BrianGreenhill/note/pkg/git"
)

func init() {
	rootCmd.AddCommand(syncCmd)
	syncCmd.Flags().BoolVarP(&diff, "diff", "d", false, "Show diff")
}

var (
	diff    bool
	syncCmd = &cobra.Command{
		Use:   "sync",
		Short: "Sync your notes",
		Run:   sync,
	}
)

func sync(cmd *cobra.Command, args []string) {
	if diff {
		fmt.Println("Displaying notes that need to be synced")
		if !git.Diff() {
			fmt.Println("Could not get diff")
		}
	} else {
		fmt.Println("Syncing notes to GitHub")
		if !(git.StashPush() && git.Pull() && git.StashPop()) {
			fmt.Println("Could not pull down repo")
		}
		if !(git.Add() && git.Commit()) {
			fmt.Println("Could not commit changes")
		}
		if !(git.Push()) {
			fmt.Println("Could not push changes to repo")
		}
	}
}

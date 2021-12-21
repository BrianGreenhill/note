package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().StringVarP(&search, "search", "s", "", "Search string for note")
}

var (
	search  string
	editCmd = &cobra.Command{
		Use:   "edit",
		Short: "Edit a note",
		Run:   editNote,
	}
)

func editNote(cmd *cobra.Command, args []string) {
	files, err := ioutil.ReadDir(viper.GetString("notedir"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	results := []string{}
	for _, f := range files {
		if f.Name() == ".git" {
			continue
		}
		// search
		if strings.Contains(f.Name(), search) {
			results = append(results, fmt.Sprintf("%s/%s", viper.GetString("notedir"), f.Name()))
		}
	}

	prompt := promptui.Select{
		Label: "Select File",
		Items: results,
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	vimCmd := exec.Command(viper.GetString("editor"), result)
	vimCmd.Stdin = os.Stdin
	vimCmd.Stdout = os.Stdout
	if err := vimCmd.Run(); err != nil {
		log.Fatal(err)
	}
}

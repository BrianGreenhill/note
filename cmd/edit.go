package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/BrianGreenhill/note/pkg/search"
)

func init() {
	rootCmd.AddCommand(editCmd)
}

var (
	editCmd = &cobra.Command{
		Use:   "edit",
		Short: "Edit a note",
		Run:   editNote,
	}
)

func editNote(cmd *cobra.Command, args []string) {
	files, err := ioutil.ReadDir(viper.GetString("notedir"))
	if err != nil {
		log.Fatal(err)
	}

	s := search.NewFileSearcher(files)

	_, result, err := s.SearchPrompt("Select Note").Run()
	if err != nil {
		log.Fatal(err)
	}

	editorCmd := exec.Command(viper.GetString("editor"), fmt.Sprintf("%s/%s", viper.GetString("notedir"), result))
	editorCmd.Stdin = os.Stdin
	editorCmd.Stdout = os.Stdout
	if err := editorCmd.Run(); err != nil {
		log.Fatal(err)
	}
}

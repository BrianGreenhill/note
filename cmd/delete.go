package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/BrianGreenhill/note/pkg/search"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var (
	deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete your notes",
		Run:   del,
	}
)

func del(cmd *cobra.Command, args []string) {
	files, err := ioutil.ReadDir(viper.GetString("notedir"))
	if err != nil {
		log.Fatal(err)
	}

	s := search.NewFileSearcher(files)

	_, result, err := s.SearchPrompt("Select Note to Delete").Run()
	if err != nil {
		log.Fatal(err)
	}

	confPrompt := promptui.Prompt{
		Label:     fmt.Sprintf("Delete %s", result),
		IsConfirm: true,
	}

	shouldDelete, err := confPrompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	if strings.ToLower(shouldDelete) == "y" {
		if err = os.Remove(fmt.Sprintf("%s/%s", viper.GetString("notedir"), result)); err != nil {
			fmt.Println("failed to delete file", result)
		}
		return
	}
}

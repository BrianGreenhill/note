package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var (
	tsformat = "2006-01-02 15:04"
	listCmd  = &cobra.Command{
		Use:   "list",
		Short: "List your notes",
		Run: func(cmd *cobra.Command, args []string) {
			files, err := ioutil.ReadDir(viper.GetString("notedir"))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println("Your Notes:")
			fmt.Println("")
			sort.Slice(files, func(i, j int) bool {
				return files[i].ModTime().After(files[j].ModTime())
			})

			for i, f := range files {
				if f.Name() == ".git" {
					continue
				}
				fmt.Println(i, f.Name(), f.ModTime().Local().Format(tsformat))
			}
		},
	}
)

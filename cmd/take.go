package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(takeCmd)
	takeCmd.Flags().StringVarP(&note, "note", "n", "", "Note to take as string")
	takeCmd.Flags().StringVar(&format, "format", "md", "File format for note")
	if err := viper.BindPFlag("format", takeCmd.Flags().Lookup("format")); err != nil {
		log.Fatal(err)
	}
}

const (
	TSFORMAT = "20060102150405"
)

var (
	note    string
	format  string
	takeCmd = &cobra.Command{
		Use:   "take",
		Short: "Take your notes",
		Args:  cobra.ExactArgs(1),
		Run:   takeNote,
	}
)

func takeNote(cmd *cobra.Command, args []string) {
	timestamp := time.Now()
	note := []byte("# " + args[0])
	filename := fmt.Sprintf("%s/%s_%s.%s", viper.GetString("notedir"), timestamp.Format(TSFORMAT), args[0], viper.GetString("format"))
	err := ioutil.WriteFile(filename, note, 0644)
	if err != nil {
		log.Fatalf("Could not write file %+v", err)
	}

	editorCmd := exec.Command(viper.GetString("editor"), filename)
	editorCmd.Stdin = os.Stdin
	editorCmd.Stdout = os.Stdout
	if err := editorCmd.Run(); err != nil {
		log.Fatal(err)
	}
}

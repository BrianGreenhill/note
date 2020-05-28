package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(takeCmd)
	takeCmd.Flags().StringVarP(&note, "note", "n", "", "Note to take as string")
	takeCmd.Flags().StringVar(&format, "format", "md", "File format for note")
	viper.BindPFlag("format", takeCmd.Flags().Lookup("format"))
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

	fmt.Printf("Taking your note: %s. Find it here: %s", args[0], filename)
}

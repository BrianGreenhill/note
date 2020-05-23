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
}

const (
	TSFORMAT = "20060102150405"
)

var (
	note    string
	takeCmd = &cobra.Command{
		Use:   "take",
		Short: "Take your notes",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			timestamp := time.Now()
			note := []byte(args[0])
			filename := fmt.Sprintf("%s/%s_note", viper.GetString("notedir"), timestamp.Format(TSFORMAT))
			err := ioutil.WriteFile(filename, note, 0644)
			if err != nil {
				log.Fatalf("Could not write file %+v", err)
			}

			fmt.Printf("Taking your note: %s. Find it here: %s", args[0], filename)
		},
	}
)

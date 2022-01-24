package cmd

import (
	"fmt"
	"log"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile  string
	noteDir  string
	gitUser  string
	gitToken string
	editor   string

	rootCmd = &cobra.Command{
		Use:   "note",
		Short: "take note of your notes",
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.note.yaml)")
	rootCmd.PersistentFlags().StringVar(&noteDir, "notedir", "", "Note directory")
	rootCmd.PersistentFlags().StringVar(&gitUser, "gituser", "", "Github user")
	rootCmd.PersistentFlags().StringVar(&gitToken, "gittoken", "", "Github token")
	rootCmd.PersistentFlags().StringVar(&editor, "editor", "", "Editor")
	if err := viper.BindPFlag("notedir", rootCmd.PersistentFlags().Lookup("notedir")); err != nil {
		log.Fatal(err)
	}
	if err := viper.BindPFlag("gituser", rootCmd.PersistentFlags().Lookup("gituser")); err != nil {
		log.Fatal(err)
	}
	if err := viper.BindPFlag("gittoken", rootCmd.PersistentFlags().Lookup("gittoken")); err != nil {
		log.Fatal(err)
	}
	if err := viper.BindPFlag("editor", rootCmd.PersistentFlags().Lookup("editor")); err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".note")
	}
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func Execute() error {
	return rootCmd.Execute()
}

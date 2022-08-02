/*
Copyright © 2022 devenes ahmedenesturan@gmail.com

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli-world",
	Short: "A basic CLI application",
	Long: `
	A longer description about how to use this application
	For example: cli-world
	For example: cli-world -o
	`,

	Run: func(cmd *cobra.Command, args []string) {
		flagVal, err := cmd.Flags().GetBool("othermessage")
		if err != nil {
			return
		}
		if flagVal {
			fmt.Println("I like these words better!")
			return
		}
		fmt.Println("Hello World!")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli-world.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("othermessage", "o", false, "Toggle the other message")
}

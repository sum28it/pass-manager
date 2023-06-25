/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/sum28it/pass-manager/pkg/auth"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if !auth.IsInit() {
			fmt.Println("app not initialized")
		}

		// Authenticate user before resetting
		err := auth.Authenticate(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		var choice string
		fmt.Printf("This will remove all your data including any password that might be saved.\nAre you sure you want to do this? (Yes/No)")
		fmt.Scanf("%s", &choice)
		choice = strings.ToUpper(choice)

		switch choice {
		case "YES":
			err := os.RemoveAll("files")
			if err != nil {
				fmt.Println("Unable to reset", err)
			}
		case "NO":
			return

		default:
			fmt.Println("Invalid Input!")
		}

	},
}

func init() {
	rootCmd.AddCommand(resetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

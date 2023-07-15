/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/sum28it/pass-manager/pkg/user"
	"golang.org/x/term"
)

// resetPassCmd represents the resetPass command
var resetSecretCmd = &cobra.Command{
	Use:   "resetSecret",
	Short: "Command to reset your secret",
	Long:  `This command is used to reset your secret. It's a safe process and all the data would remain unaffected.`,
	Run: func(cmd *cobra.Command, args []string) {

		if !user.IsInit() {
			fmt.Println("App not initialized!\nUse init command to initialize the application first.")
		}

		fmt.Print("Your secret: ")
		secret, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Error reading secret")
			return
		}
		fmt.Println()

		fmt.Print("Your secret: ")
		newSecret, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Error reading secret")
			return
		}
		fmt.Println()

		fmt.Print("Your secret: ")
		reEnteredSecret, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Error reading secret")
			return
		}
		fmt.Println()

		if string(reEnteredSecret) != string(newSecret) {
			fmt.Println("New Secrets do not match!")
			return
		}

		err = user.ResetSecret(string(secret), string(newSecret))

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Your secret has been succesfully changed")

	},
}

func init() {
	rootCmd.AddCommand(resetSecretCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resetPassCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resetPassCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

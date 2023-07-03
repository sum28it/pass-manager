/*
Copyright © 2023 sum28it prasad28sumit@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/sum28it/pass-manager/pkg/user"
	"golang.org/x/term"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Resets the application",
	Long: `This command resets the application and deletes all the app data. It takes the secret as an argument.
WARNING: [All your data will be lost]`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Print("Your secret: ")
		secret, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Error reading secret")
			return
		}
		fmt.Println()
		var choice string
		fmt.Println("This will remove all your data including any password that might be saved.\nAre you sure you want to do this? (Yes/No)")
		fmt.Scanf("%s", &choice)
		choice = strings.ToUpper(choice)

		switch choice {
		case "YES":
			err := user.Reset(string(secret))
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Your application has been successsfully reset.\nUse init command again before adding users.")
		default:
			fmt.Println("Your data is safe!")
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

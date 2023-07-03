/*
Copyright © 2023 sum28it prasad28sumit@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/sum28it/pass-manager/pkg/user"
	"golang.org/x/term"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Used for adding a new user ",
	Long: `This command is used to add a new user data. It has multiple flags, some of which 
are required and others are optional. This command takes a mandatory argument i.e the secret for
the application`,
	Run: func(cmd *cobra.Command, args []string) {

		// Read password and secret from the terminal
		fmt.Print("Enter password: ")
		password, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Error reading password")
			return
		}
		fmt.Print("\nYour secret: ")
		secret, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Error reading secret")
			return
		}
		fmt.Println()

		u := user.User{
			App:         cmd.Flag("app").Value.String(),
			UserId:      cmd.Flag("userId").Value.String(),
			Email:       cmd.Flag("email").Value.String(),
			Password:    string(password),
			ModifiedAt:  time.Now().Format("2006-01-02 15:04:05"),
			Description: cmd.Flag("description").Value.String(),
		}
		users, err := user.Add(u, string(secret))
		if err != nil {
			fmt.Println(err)
			for _, u := range users {
				fmt.Println(u.Print())
			}
			return
		}

		fmt.Println("User Added!")
		fmt.Println(u.PrintLong())
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().StringP("app", "a", "", "App name (required)")
	addCmd.Flags().StringP("userId", "u", "", "User ID")
	addCmd.Flags().StringP("email", "e", "", "Email")
	addCmd.Flags().StringP("description", "d", "", "Description")

	addCmd.MarkFlagRequired("app")

}

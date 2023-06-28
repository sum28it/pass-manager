/*
Copyright © 2023 sum28it prasad28sumit@gmail.com
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/sum28it/pass-manager/pkg/user"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes one or more user data",
	Long: `This command is for deleting one or more app data. It takes an argument secret and has flags for specifying 
	the app data to be deleted`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		u := user.User{
			App:    cmd.Flag("app").Value.String(),
			Email:  cmd.Flag("email").Value.String(),
			UserId: cmd.Flag("userId").Value.String(),
		}
		var users []user.User
		users, err := user.Delete(u, args[0], false)
		if err != nil {
			fmt.Println(err)
			return
		}
		if len(users) > 1 {
			var choice string
			for _, u := range users {
				fmt.Println(u.Print())
			}
			fmt.Println("More than one such user exists", "Do you want to delete all(Yes/No)?")
			fmt.Scanf("%s", &choice)
			choice = strings.ToUpper(choice)
			if choice == "YES" {
				user.Delete(u, args[0], true)
			}
		}
		fmt.Println("Deleted!")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	deleteCmd.Flags().StringP("app", "a", "", "App name")
	deleteCmd.Flags().StringP("email", "e", "", "Email")
	deleteCmd.Flags().StringP("userId", "u", "", "User ID")
}

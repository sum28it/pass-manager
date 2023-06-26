/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sum28it/pass-manager/pkg/user"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve user data",
	Long: `Thsi command is used to retrieve usert data. It takes two arguments. 
	First is the name of the app you want to get and second is your secret.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// value, _ := cmd.Flags().GetBool("long")
		// fmt.Println("get called", value, args)

		u := &user.User{
			App:    cmd.Flag("app").Value.String(),
			Email:  cmd.Flag("email").Value.String(),
			UserId: cmd.Flag("userId").Value.String(),
		}
		u, err := user.Get(u, args[0])

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(*u)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	getCmd.Flags().BoolP("long", "l", false, "Print user info in long format")
	getCmd.Flags().StringP("app", "a", "", "App name")
	getCmd.Flags().StringP("email", "e", "", "Email")
	getCmd.Flags().StringP("userId", "u", "", "User ID")

}

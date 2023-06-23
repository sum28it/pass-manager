/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sum28it/pass-manager/user"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		u := &user.User{
			App:         cmd.Flag("app").Value.String(),
			UserId:      cmd.Flag("userId").Value.String(),
			Email:       cmd.Flag("email").Value.String(),
			Password:    cmd.Flag("password").Value.String(),
			Description: cmd.Flag("description").Value.String(),
		}
		err := user.Add(u)
		if err != nil {
			fmt.Println(err)
		}
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
	addCmd.Flags().StringP("password", "p", "", "Password")
	addCmd.Flags().StringP("description", "d", "", "Description")

	addCmd.MarkFlagRequired("app")
	addCmd.MarkFlagRequired("password")

}

/*
Copyright © 2023 sum28it prasad28sumit@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/sum28it/pass-manager/pkg/user"
	"golang.org/x/term"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve user data",
	Long: `This command is used to retrieve app data. It takes an argument secret and flags for specifying the 
details of queried app.`,
	Run: func(cmd *cobra.Command, args []string) {
		// value, _ := cmd.Flags().GetBool("long")
		// fmt.Println("get called", value, args)

		fmt.Print("Your secret: ")
		secret, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Error reading secret")
			return
		}
		fmt.Println()
		u := user.User{
			App:    cmd.Flag("app").Value.String(),
			Email:  cmd.Flag("email").Value.String(),
			UserId: cmd.Flag("userId").Value.String(),
		}
		result, err := user.Get(u, string(secret))

		if err != nil {
			fmt.Println(err)
			return
		}
		verbose, _ := cmd.Flags().GetBool("verbose")
		for _, u := range result {
			if verbose {
				fmt.Println(u.PrintLong())
			} else {
				fmt.Println(u.Print())
			}
		}
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

	getCmd.Flags().BoolP("verbose", "v", false, "Print user info in verbose mode")
	getCmd.Flags().StringP("app", "a", "", "Get App")
	getCmd.Flags().StringP("email", "e", "", "Email get")
	getCmd.Flags().StringP("userId", "u", "", "User ID")

}

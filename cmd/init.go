/*
Copyright © 2023 sum28it prasad28sumit@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sum28it/pass-manager/pkg/user"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes the application",
	Long: `init command is the first command to be run after installing the application.
	It takes an argument which is the secret for accessing the appication.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		dir, err := user.Init(args[0])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Data is stored at: ", dir)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

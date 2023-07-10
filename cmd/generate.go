/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/sum28it/pass-manager/pkg/auth"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		rand.NewSource(time.Now().Unix())
		var err error
		var passLen int

		len := cmd.Flag("length").Value.String()
		if len != "" {
			passLen, err = strconv.Atoi(len)
		} else {
			passLen = 8 + rand.Intn(4)
		}
		if err != nil || passLen <= 0 {
			fmt.Println("ERROR: invalid length value")
			return
		}

		alpha, _ := cmd.Flags().GetBool("alpha")
		num, _ := cmd.Flags().GetBool("num")
		alphaNum, _ := cmd.Flags().GetBool("alpha-num")

		if (alpha && num) || (alpha && alphaNum) || (num && alphaNum) {
			fmt.Println("ERROR: multiple flags passed for type of password")
		}

		var password string
		switch {
		case alpha:
			password = auth.GenerateRandomPassword(passLen, auth.ALPHA)
		case num:
			password = auth.GenerateRandomPassword(passLen, auth.NUM)
		case alphaNum:
			password = auth.GenerateRandomPassword(passLen, auth.ALPHA_NUM)
		default:
			password = auth.GenerateRandomPassword(passLen, auth.ALPHA_NUM_SPECIAL)
		}

		fmt.Println(password)

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	generateCmd.Flags().StringP("length", "l", "", "Length of the password")
	generateCmd.Flags().BoolP("alpha", "a", false, "Only alphabetic characters")
	generateCmd.Flags().BoolP("alpha-num", "A", false, "Only alpha-numeric characters")
	generateCmd.Flags().BoolP("num", "n", false, "Only numeric characters")
}

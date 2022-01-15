/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// yieldCmd represents the yield command
var yieldCmd = &cobra.Command{
	Use:   "yield",
	Short: "Gives you a random problem to solve today.",
	Long: `This command yields a problem for you to solve today.
This command takes into account, your history with leetcode (only using a local store),
and then yields a random problem that you should try to solve now.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("yield called")
	},
}

func init() {
	rootCmd.AddCommand(yieldCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// yieldCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// yieldCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

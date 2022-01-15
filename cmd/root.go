/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "leetgo",
	Short: "leetgo is a CLI application to help you study for DSA interviews using Leetcode",
	Long: `leetgo uses spaced repetition to help you decide what to study next for leetcode.
It uses a corpus of leetcode problems to help you study. It can be used to record your
efforts, note how long you took to solve something, how many times you've tried solving a problem,
and accordingly gives you redos or tasks you with revision depending on how
long it has been since you've tackled a problem before.

Leetgo is named after John Lithgow, for the pun.

> Time sneaks up on you like a windshield on a bug.
- **John Lithgow** `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.leetgo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}



/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "mogi",
		Short: "A fast and simple way to create mock data",
		Long: `A fast and simple way to generate mock data for 
various platforms. The plan for now is to
support platforms that I use at work or might use for 
personal projects, such as Shopify and Smaregi as well as
general SQL scripts. 
	
Eventually, I would like to support other platforms, the Shopfiy
and Smaregi APIs, and more!
	
This is also just a pet project for me to learn Go with,
so I'm not actually expecting anyone to use it except myself. ☺️`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mogi.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of mogi",
	Long:  `Is there really anything more I can say about this?`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mogi - mock data cli v0.1.0")
	},
}

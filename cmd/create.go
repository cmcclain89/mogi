package cmd

import (
	"fmt"
	"mogi/shopify"

	"github.com/spf13/cobra"
)

var Format string

func init() {
	createCmd.Flags().StringVarP(&Format, "format", "f", "", "Format of the mock data")
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create mock data.",
	Long: `Generate mock data. Pass in platform as args, format as flags. CSV is the default.

	Current supported args:
shopify
`,
	ValidArgs: []string{"shopify", "smaregi"},
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			execute_for_arg(arg)
		}

		fmt.Println("neat")
	},
}

func execute_for_arg(arg string) {
	switch arg {
	case "shopify":
		fmt.Printf("Generating mock data for %s.\r\n", arg)
		shopify.GenerateCSV()
	default:
		fmt.Printf("%s isn't an accepted argument, ignoring.\r\n", arg)
	}
}

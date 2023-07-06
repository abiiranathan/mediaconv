package cmd

import (
	"github.com/spf13/cobra"
)

// cutVideoCmd represents the cutVideo command
var cutVideoCmd = &cobra.Command{
	Use:   "cutVideo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return ExtractVideoSegment(input, output, from, to)
	},
}

func init() {
	rootCmd.AddCommand(cutVideoCmd)

	cutVideoCmd.Flags().StringVarP(&from, "start", "s", "", "Starting timestamp")
	cutVideoCmd.Flags().StringVarP(&to, "end", "e", "", "Ending timestamp")

	cutVideoCmd.MarkFlagRequired("start")
	cutVideoCmd.MarkFlagRequired("end")
}

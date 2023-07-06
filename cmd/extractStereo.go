package cmd

import (
	"github.com/spf13/cobra"
)

var outputRightFilename string

// extractStereoCmd represents the extractStereo command
var extractStereoCmd = &cobra.Command{
	Use:   "extractStereo",
	Short: "Seperate left and right stereo channels from input",
	RunE: func(cmd *cobra.Command, args []string) error {
		return ExtractStereoChannels(input, output, outputRightFilename)
	},
}

func init() {
	rootCmd.AddCommand(extractStereoCmd)
	extractStereoCmd.Flags().StringVarP(&outputRightFilename, "right", "r", "", "Path to save the right stereo")
	extractStereoCmd.MarkFlagRequired("right")
}

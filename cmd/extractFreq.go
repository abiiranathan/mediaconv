package cmd

import (
	"github.com/spf13/cobra"
)

var (
	startFreq, endFreq int
)

// extractFreqCmd represents the extractFreq command
var extractFreqCmd = &cobra.Command{
	Use:   "extractFreq",
	Short: "Generate a new audio file with frequencies within the specified range",
	RunE: func(cmd *cobra.Command, args []string) error {
		return ExtractAudioFrequencies(input, output, startFreq, endFreq)
	},
}

func init() {
	rootCmd.AddCommand(extractFreqCmd)

	extractFreqCmd.Flags().IntVarP(&startFreq, "low", "s", 0, "Lowest frequency")
	extractFreqCmd.Flags().IntVarP(&endFreq, "high", "u", 0, "Highest frequency")

	extractFreqCmd.MarkFlagRequired("low")
	extractFreqCmd.MarkFlagRequired("high")
}

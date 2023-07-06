package cmd

import (
	"github.com/spf13/cobra"
)

var silenceThreshold float64

// trimSilenceCmd represents the trimSilence command
var trimSilenceCmd = &cobra.Command{
	Use:   "trimSilence",
	Short: "Remove silent portions of audio from video/audio",
	RunE: func(cmd *cobra.Command, args []string) error {
		return TrimSilence(input, output, silenceThreshold)
	},
}

func init() {
	rootCmd.AddCommand(trimSilenceCmd)
	trimSilenceCmd.Flags().Float64VarP(&silenceThreshold, "threshold", "t", -30.0,
		"Silence threshold: Specifies the silence threshold in decibels (dB). Any audio portion below this threshold will be considered as silence and removed")
}

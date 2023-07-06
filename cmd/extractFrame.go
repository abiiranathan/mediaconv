package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

var (
	timestampString string
)

// extractFrameCmd represents the resize command
var extractFrameCmd = &cobra.Command{
	Use:   "extractFrame",
	Short: "Extract video frame",
	Long: `
Extract a video frame for each of a comma-seperated string of timestamps.e.g 
12:30,01:30:12,00:00:07
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		timestamps := []string{}
		for _, s := range strings.Split(timestampString, ",") {
			timestamps = append(timestamps, strings.TrimSpace(s))
		}
		return ExtractFrameAtTimestamps(input, output, timestamps)
	},
}

func init() {
	rootCmd.AddCommand(extractFrameCmd)
	extractFrameCmd.Flags().StringVarP(&timestampString, "timestamps", "t", "", "comma-seperated timestamps")
	extractFrameCmd.MarkFlagRequired("timestamps")
}

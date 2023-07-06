package cmd

import (
	"github.com/spf13/cobra"
)

var (
	codec          string
	startTimestamp string
	endTimestamp   string
)

// extractAudioCmd represents the resize command
var extractAudioCmd = &cobra.Command{
	Use:   "extractAudio",
	Short: "Extract audio from a video",
	Long: `
Extract audio from a video. Optional timestamps allow you to trim the output.
If codec is not provided, it will be copied from source audio stream.
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return ExtractAudioFromVideoWithTimestamps(input, output, codec, startTimestamp, endTimestamp)
	},
}

func init() {
	rootCmd.AddCommand(extractAudioCmd)
	extractAudioCmd.Flags().StringVarP(&codec, "codec", "c", "copy", "Output codec")
	extractAudioCmd.Flags().StringVarP(&startTimestamp, "start", "s", "", "Timestamp to start cut.")
	extractAudioCmd.Flags().StringVarP(&endTimestamp, "end", "e", "", "Timestamp to stop cut.")
}

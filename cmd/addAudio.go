package cmd

import (
	"github.com/spf13/cobra"
)

var audioInput string

// addAudioCmd represents the addAudio command
var addAudioCmd = &cobra.Command{
	Use:   "addAudio",
	Short: "Add audio to a video. Trim at the shortest",
	RunE: func(cmd *cobra.Command, args []string) error {
		return MergeAudioWithVideo(audioInput, input, output)
	},
}

func init() {
	rootCmd.AddCommand(addAudioCmd)

	addAudioCmd.Flags().StringVarP(&input, "input", "i", "", "Video Input")
	addAudioCmd.Flags().StringVarP(&audioInput, "audio", "a", "", "Audio Input")

	addAudioCmd.MarkFlagRequired("input")
	addAudioCmd.MarkFlagRequired("audio")

}

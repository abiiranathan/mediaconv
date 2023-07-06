package cmd

import (
	"github.com/spf13/cobra"
)

var subtitleFilename string

// addSubtitlesCmd represents the addSubtitles command
var addSubtitlesCmd = &cobra.Command{
	Use:   "addSubtitles",
	Short: "Add subtitles to video from a file",
	RunE: func(cmd *cobra.Command, args []string) error {
		return AddSubtitles(input, output, subtitleFilename)
	},
}

func init() {
	rootCmd.AddCommand(addSubtitlesCmd)
	addSubtitlesCmd.Flags().StringVarP(&subtitleFilename, "subtitles", "s", "", "subtitles filename")
	addSubtitlesCmd.MarkFlagRequired("subtitles")
}

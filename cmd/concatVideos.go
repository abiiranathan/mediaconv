package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var inputFileNames []string

// concatVideosCmd represents the concatVideos command
var concatVideosCmd = &cobra.Command{
	Use:   "concatVideos",
	Short: "Concatenate multiple videos",
	Long: `Join multiple videos into a single video stream.

	Because video may have different dimensions, you must specify
	width and height to properly resize the video.
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(inputFileNames) < 2 {
			return fmt.Errorf("pass 2 or more input videos to concatenate")
		}
		return ConcatenateVideos(inputFileNames, output, width, height)
	},
}

func init() {
	rootCmd.AddCommand(concatVideosCmd)
	concatVideosCmd.Flags().StringSliceVarP(&inputFileNames, "input", "i", []string{}, "Input file names")
	concatVideosCmd.Flags().IntVarP(&width, "width", "W", 0, "Final video width")
	concatVideosCmd.Flags().IntVarP(&height, "height", "H", 0, "Final video height")
	concatVideosCmd.MarkFlagRequired("input")
	concatVideosCmd.MarkFlagRequired("width")
	concatVideosCmd.MarkFlagRequired("height")
}

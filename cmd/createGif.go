package cmd

import (
	"github.com/spf13/cobra"
)

var fps int
var from, to string

// createGifCmd represents the createGif command
var createGifCmd = &cobra.Command{
	Use:   "createGif",
	Short: "Create an animated GIF from a video",
	RunE: func(cmd *cobra.Command, args []string) error {
		return ConvertToGif(input, output, width, height, fps, from, to)
	},
}

func init() {
	rootCmd.AddCommand(createGifCmd)
	createGifCmd.Flags().IntVarP(&width, "width", "W", 360, "Width of the GIF")
	createGifCmd.Flags().IntVarP(&height, "height", "H", 240, "Height of the GIF")
	createGifCmd.Flags().IntVarP(&fps, "fps", "f", 5, "Frames per second")
	createGifCmd.Flags().StringVarP(&from, "start", "s", "", "Starting timestamp")
	createGifCmd.Flags().StringVarP(&to, "end", "e", "", "Ending timestamp")

	createGifCmd.MarkFlagRequired("start")
	createGifCmd.MarkFlagRequired("end")

}

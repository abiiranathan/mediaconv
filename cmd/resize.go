package cmd

import (
	"github.com/spf13/cobra"
)

var (
	width  int // Width of output video/image
	height int // Height of output video/image
)

// resizeCmd represents the resize command
var resizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "Resize Video",
	Long:  `Resize a video by specifying the width and height in pixels`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return ResizeVideo(input, output, width, height)
	},
}

func init() {
	rootCmd.AddCommand(resizeCmd)
	resizeCmd.Flags().IntVarP(&width, "width", "W", 720, "Width of the output video")
	resizeCmd.Flags().IntVarP(&height, "height", "H", 480, "Height to the output video")

}

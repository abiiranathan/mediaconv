package cmd

import (
	"github.com/spf13/cobra"
)

var (
	image    string
	position string
	wmWidth  int
	wmHeight int
)

const (
	OVERLAY_CENTER        = "(main_w-overlay_w)/2:(main_h-overlay_h)/2"
	OVERLAY_TOP_LEFT      = "10:10"
	OVERLAY_TOP_RIGHT     = "(main_w-overlay_w)-10:10"
	OVERLAY_TOP_CENTER    = "(main_w-overlay_w)/2:10"
	OVERLAY_BOTTOM_LEFT   = "10:(main_h-overlay_h)-10"
	OVERLAY_BOTTOM_RIGHT  = "(main_w-overlay_w)-10:(main_h-overlay_h)-10"
	OVERLAY_BOTTOM_CENTER = "(main_w-overlay_w)/2:(main_h-overlay_h)-10"
)

var positions = map[string]string{
	"c":  OVERLAY_CENTER,
	"bl": OVERLAY_BOTTOM_LEFT,
	"br": OVERLAY_BOTTOM_RIGHT,
	"bc": OVERLAY_BOTTOM_CENTER,
	"tl": OVERLAY_TOP_LEFT,
	"tc": OVERLAY_TOP_CENTER,
	"tr": OVERLAY_TOP_RIGHT,
}

// imageWatermarkCmd represents the textWatermark command
var imageWatermarkCmd = &cobra.Command{
	Use:   "imageWatermark",
	Short: "Apply a image watermark to video",
	Long: `

Valid overlay positions:
- "c" for Center
- "bl" for bottom left
- "br" for bottom right
- "bc" for bottom center
- "tl" for top left
- "tr" for top right
- "tc" for top center
- Any other valid ffmpeg overlay filter e.g 10:10. If it's not valid, expect an ffmpeg error.
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var validPos string
		if pos, ok := positions[position]; ok {
			validPos = pos
		} else {
			// Pass position as is
			validPos = position
		}
		return ApplyImageWatermark(input, output, image, validPos, wmWidth, wmWidth)
	},
}

func init() {
	rootCmd.AddCommand(imageWatermarkCmd)

	imageWatermarkCmd.Flags().StringVarP(&image, "image", "I", "", "Image to use as a watermark")
	imageWatermarkCmd.Flags().StringVarP(&position, "position", "p", OVERLAY_BOTTOM_RIGHT, "Position to apply the watermark")
	imageWatermarkCmd.Flags().IntVarP(&wmWidth, "width", "W", 100, "Watermark image width")
	imageWatermarkCmd.Flags().IntVarP(&wmHeight, "height", "H", 100, "Watermark image height")

	imageWatermarkCmd.MarkFlagRequired("image")

}

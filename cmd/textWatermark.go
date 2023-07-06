package cmd

import (
	"github.com/spf13/cobra"
)

var (
	text     string
	fontSize float32
	fontPath string
	fgColor  string
	bgColor  string
)

// textWatermarkCmd represents the textWatermark command
var textWatermarkCmd = &cobra.Command{
	Use:   "textWatermark",
	Short: "Apply a text watermark to video",
	Long: `
Here are some examples of valid color strings in the format "0xRRGGBBAA":

Fully opaque red: "0xFF0000FF"
Semi-transparent blue: "0x0000FF80" (80 represents 50% opacity)
Fully opaque green: "0x00FF00FF"
Semi-transparent white: "0xFFFFFF80" (80 represents 50% opacity)
Fully opaque black: "0x000000FF"
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return ApplyTextWatermark(input, output, text, position, bgColor, fgColor, fontSize, fontPath)
	},
}

func init() {
	rootCmd.AddCommand(textWatermarkCmd)

	textWatermarkCmd.Flags().StringVarP(&text, "text", "t", "", "Text to use as a watermark")
	textWatermarkCmd.Flags().StringVarP(&position, "position", "p", "center", "Position to apply the watermark")
	textWatermarkCmd.Flags().StringVarP(&fontPath, "fontfile", "f", "", "Font file path to use for the watermark text")
	textWatermarkCmd.Flags().Float32VarP(&fontSize, "fontsize", "s", 16, "Font size for the watermark text")
	textWatermarkCmd.Flags().StringVarP(&fgColor, "fg", "c", "0xFFFFFFFF", "Foreground color")
	textWatermarkCmd.Flags().StringVarP(&bgColor, "bg", "b", "0x000000FF", "Background color")
	textWatermarkCmd.MarkFlagRequired("text")
}

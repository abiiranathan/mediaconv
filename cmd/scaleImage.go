/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var (
	aspectRatio float32
)

// scaleImageCmd represents the scaleImage command
var scaleImageCmd = &cobra.Command{
	Use:   "scaleImage",
	Short: "Scale image to specified width, height or aspect ratio",
	Long:  `You must specify either width & height or aspect ratio`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if width > 0 && height > 0 {
			return ScaleImage(input, output, width, height)
		}
		return ScaleImageToAspectRation(input, output, aspectRatio)
	},
}

func init() {
	rootCmd.AddCommand(scaleImageCmd)

	scaleImageCmd.Flags().Float32VarP(&aspectRatio, "aspect", "a", 0.5, "Aspect ratio")
	scaleImageCmd.Flags().IntVarP(&width, "width", "W", 0, "Width")
	scaleImageCmd.Flags().IntVarP(&height, "height", "H", 0, "Height")

	scaleImageCmd.MarkFlagsRequiredTogether("width", "height")
	scaleImageCmd.MarkFlagsMutuallyExclusive("aspect", "width")
	scaleImageCmd.MarkFlagsMutuallyExclusive("aspect", "height")
}

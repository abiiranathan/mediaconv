/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// img2VideoCmd represents the img2Video command
var img2VideoCmd = &cobra.Command{
	Use:   "img2Video",
	Short: "Create a video from a folder of images",
	Long:  `If the width and height are zero, we use the default image dimensions.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return ImagesToVideo(input, output, fps, width, height)
	},
}

func init() {
	rootCmd.AddCommand(img2VideoCmd)

	// Override input from global persistent flags
	img2VideoCmd.Flags().StringVarP(&input, "input", "i", "", "Input format")
	img2VideoCmd.Flags().IntVarP(&width, "width", "W", 0, "Width of the target video")
	img2VideoCmd.Flags().IntVarP(&height, "height", "H", 0, "Height of the target video")
	img2VideoCmd.Flags().IntVarP(&fps, "fps", "f", 30, "Frames per second")

	img2VideoCmd.MarkFlagRequired("input")
}

package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
)

var outputQuality int
var outputDir string

// convertVideoCmd represents the convertVideo command
var convertVideoCmd = &cobra.Command{
	Use:   "convertVideo",
	Short: "Convert one or more videos",
	Long: `
Input:  globPattern for the input videos. e.g ./*.mkv
Output: File extension for the output files e.g mp4

The glob pattern is processed and expanded and each video converted in parallel.
You can specify the quality of the video.
Quality value should be between 1 and 31, where 1 represents the highest quality.
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		outputDirectory, err := filepath.Abs(outputDir)
		if err != nil {
			return err
		}
		return ConvertVideos(input, output, outputDirectory, outputQuality)
	},
}

func init() {
	rootCmd.AddCommand(convertVideoCmd)
	convertVideoCmd.Flags().IntVarP(&outputQuality, "quality", "q", 23, "Video quality of converted videos")
	convertVideoCmd.Flags().StringVarP(&outputDir, "dir", "d", "", "Output directory. Default to same dir as source")
}

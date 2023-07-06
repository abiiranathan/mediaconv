package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var imageFormat string

// genImagesCmd represents the genImages command
var genImagesCmd = &cobra.Command{
	Use:   "genImages",
	Short: "Generate Images from video",
	Long: `
Customize images by specifying custom dimensions, output format,
timestamps where to start and/or end from.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dirname, err := filepath.Abs(output)
		if err != nil {
			return err
		}

		// Create all intermediate directories if not exists.
		err = os.MkdirAll(dirname, 0755)
		if err != nil {
			return err
		}
		return VideoToImages(input, output, imageFormat, width, height, from, to)
	},
}

func init() {
	rootCmd.AddCommand(genImagesCmd)

	genImagesCmd.Flags().StringVarP(&imageFormat, "format", "f", "image_%003d.jpg", "Image format")
	genImagesCmd.Flags().IntVarP(&width, "width", "W", 0, "Target width of images")
	genImagesCmd.Flags().IntVarP(&height, "height", "H", 0, "Target height of images")
	genImagesCmd.Flags().StringVarP(&from, "start", "s", "", "Start timestamp")
	genImagesCmd.Flags().StringVarP(&to, "end", "e", "", "Ending timestamp")
}

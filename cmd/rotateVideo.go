package cmd

import (
	"github.com/spf13/cobra"
)

var angle int

// rotateVideoCmd represents the rotateVideo command
var rotateVideoCmd = &cobra.Command{
	Use:   "rotateVideo",
	Short: "Rotate Video",
	RunE: func(cmd *cobra.Command, args []string) error {
		return RotateVideo(input, output, angle)
	},
}

func init() {
	rootCmd.AddCommand(rotateVideoCmd)
	rotateVideoCmd.Flags().IntVarP(&angle, "angle", "a", 180, "Angle of rotation")
}

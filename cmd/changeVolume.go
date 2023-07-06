package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var volume float32

// changeVolumeCmd represents the changeVolume command
var changeVolumeCmd = &cobra.Command{
	Use:   "changeVolume",
	Short: "Adjust audio volume of audio or video",
	RunE: func(cmd *cobra.Command, args []string) error {
		if volume < 0.0 || volume > 1.0 {
			return errors.New("volume must be between 0 [muted] and 1 [Loudest]")
		}
		return ChangeAudioVolume(input, output, volume)
	},
}

func init() {
	rootCmd.AddCommand(changeVolumeCmd)
	changeVolumeCmd.Flags().Float32VarP(&volume, "volume", "v", 1.0, "Adjust volume [0.0 - 1.0]")
}

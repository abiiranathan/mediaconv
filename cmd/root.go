/*
Copyright © 2023 Abiira Nathan <nabiira2by2@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	input  string // Path to input video
	output string // Path to output video
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mediaconv",
	Short: "Simple Media Converter",
	Long: `
mediaconv

Convert videos, audio, watermark your videos,
manipulate audio using this simple cli.

Note that you must have ffmpeg installed on your machine.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Input and output flags must be passed for all subcommands and are required
	rootCmd.PersistentFlags().StringVarP(&input, "input", "i", "", "Path to the input video")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Path to the output video")

	rootCmd.MarkPersistentFlagRequired("input")
	rootCmd.MarkPersistentFlagRequired("output")
}

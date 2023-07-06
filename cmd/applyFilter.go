package cmd

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
)

var filterExpression string

// applyFilterCmd represents the applyFilter command
var applyFilterCmd = &cobra.Command{
	Use:   "applyFilter",
	Short: "Apply video filter",
	Long:  `Apply a video filter for a given ffmpeg filter expression`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if strings.TrimSpace(filterExpression) == "" {
			return errors.New("filter expression must not be empty")
		}
		return ApplyVideoFilter(input, output, filterExpression)
	},
}

func init() {
	rootCmd.AddCommand(applyFilterCmd)
	applyFilterCmd.Flags().StringVarP(&filterExpression, "filter", "f", "", "ffmpeg filter expression")
	applyFilterCmd.MarkFlagRequired("filter")
}

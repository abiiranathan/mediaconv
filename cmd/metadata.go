package cmd

import (
	"github.com/spf13/cobra"
)

// metadataCmd represents the metadata command
var metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Print media metadata",
	RunE: func(cmd *cobra.Command, args []string) error {
		return PrintMediaMetadata(input)
	},
}

func init() {
	rootCmd.AddCommand(metadataCmd)
	metadataCmd.Flags().StringVarP(&output, "output", "o", "", "Override global persistent flag.Does not do anything")
}

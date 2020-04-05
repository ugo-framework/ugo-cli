package cli

import (
	"github.com/spf13/cobra"
)

var (
	Verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "ugocli",``
	Short: "UgoCLI is a very fast golang project cli.",
	Long: `A Fast and Flexible Project CLI built on golang
          With UgoCLI everything is simple and efficient while maintaining
          speed and performance`,
	ArgAliases: []string{"ugc"},
	Run: func(cmd *cobra.Command, args []string) {
			flags, _ := cmd.Flags().GetString("version")

	},
}

func init() {
	rootCmd.PersistentFlags().StringP("version", "v", "", "Version of the UgoCLI")
}

func Execute() error {
	return rootCmd.Execute()
}

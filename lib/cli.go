package cli

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ugocli",
	Short: "UgoCLI is a very fast golang project cli.",
	Long: `A Fast and Flexible Project CLI built on golang.
With UgoCLI everything is simple and efficient while maintaining speed and performance`,
	ArgAliases: []string{"ugc"},
	Run: func(cmd *cobra.Command, args []string) {
		flags, _ := cmd.Flags().GetBool("version")
		if flags {
			fmt.Println("UGO CLI Version 0.1.0")
			fmt.Println("Go Version: ", runtime.Version())
		}

	},
}

func init() {
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Version of the UgoCLI")
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(runCmd)
}

func Execute() error {
	return rootCmd.Execute()
}

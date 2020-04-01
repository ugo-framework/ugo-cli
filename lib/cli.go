package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return err
	}
	return nil
}

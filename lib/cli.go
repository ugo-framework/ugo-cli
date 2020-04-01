package cli

import (
  "fmt"
  "github.com/spf13/cobra"
  "os"
)

var (
  Verbose bool
)

var rootCmd = &cobra.Command{
  Use:   "hugo",
  Short: "Hugo is a very fast static site generator",
  Long: `A Fast and Flexible Static Site Generator built with
        love by spf13 and friends in Go.
        Complete documentation is available at http://hugo.spf13.com`,
}

func init() {
  rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

func Execute() error {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
    return err
  }
  return nil
}

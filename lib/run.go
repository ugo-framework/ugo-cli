package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "Runs Project with live reload",
	Long:    "Runs any project with live Reload. Works best with Ugo Framework Project",
	Aliases: []string{"r"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("Project Name need to be Supplied.")
		}
		if len(args) > 1 {
			return fmt.Errorf("Too many arguments")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		fmt.Println("Project to run :", projectName)
	},
}

func init() {
	runCmd.SetUsageTemplate("\nUsage: \n\tugocli run [PROJECT_NAME]\n\tugc r [PROJECT_NAME]")
}

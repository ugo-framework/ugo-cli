package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Creates a new project",
	Long:    "Creates a new project with UGO Framework and Ugo Watcher",
	Aliases: []string{"c"},
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
		fmt.Println(projectName)
	},
}

func init() {
	createCmd.SetUsageTemplate("\nUsage: \n\tugocli create [PROJECT_NAME]\n\tugc c [PROJECT_NAME]")
}

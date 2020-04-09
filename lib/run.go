package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ugo-framework/ugo-cli/lib/runner"
	spectator "github.com/ugo-framework/ugo-spectator/lib"
	"os"
	"regexp"
)

var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "Runs Project with live reload",
	Long:    "Runs any project with live Reload. Works best with Ugo Framework Project",
	Aliases: []string{"r"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("File Name need to be Supplied to run")
		}
		if len(args) > 1 {
			return fmt.Errorf("Too many arguments")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		currDir, _ := os.Getwd()
		r, _ := regexp.Compile("^.*.(go)$")

		if !(r.MatchString(projectName)) {
			fmt.Printf("Currently only Go File is supported. Enter %s.go\n", projectName)
		} else {
			r := runner.Runner{
				Dir:      currDir,
				Filename: projectName,
			}
			watcher, err := spectator.Init(".")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			r.Run()
			for {
				select {
				case _ = <-watcher.Ch:
					r.Run()
				}
			}
		}
	},
}

func init() {
	runCmd.SetUsageTemplate("\nUsage: \n\tugocli run [FILE_NAME]\n\tugc r [FILE_NAME]")
}

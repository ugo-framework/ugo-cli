package runner

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

type Runner struct {
	Dir      string
	Filename string
}

func (r *Runner) Run() {
	fmt.Println("Running App")
	if r.Dir != "" && r.Filename != "" {
		cmd := exec.Command("go", "run", r.Dir+"/"+r.Filename)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		stderr, err := cmd.StderrPipe()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		go func() {
			merged := io.MultiReader(stdout, stderr)
			scanner := bufio.NewScanner(merged)
			for scanner.Scan() {
				msg := scanner.Text()
				fmt.Println(msg)
			}
		}()
		if err := cmd.Start(); err != nil {
			fmt.Println(err)
			return
		}
		//cmd.Wait()
	} else {
		fmt.Println("Error: Filename and Directory must be set")
	}
}

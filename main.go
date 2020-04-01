package main

import cli "github.com/ugo-framework/ugo-cli/lib"

func main() {
	err := cli.Execute()

	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"os"

	cmd "github.com/choigonyok/goopt/cmd/acvctl"
)

func main() {
	if err := cmd.ConfigAcvctl(); err != nil {
		fmt.Println((err.Error()))
		os.Exit(1)
	}

	cmd.Execute()

	if err := cmd.Execute(); err != nil {
		fmt.Println((err.Error()))
		os.Exit(1)
	}
}

// export PATH=$PATH:$(pwd)

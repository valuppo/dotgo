package main

import (
	"fmt"
	"os"
	"github.com/valuppo/dotgo/generator"
)

func main() {
	var command string

	//check if there are any os arguments
	if len(os.Args) <= 1 {
		fmt.Println("Help Dotgo Commands")
		return
	}

	command = os.Args[1]

	switch command {
	case "init":
		if len(os.Args) == 2 {
			fmt.Println("Help Init Commands")
			return
		}
		if len(os.Args) >= 4 {
			generator.Init(os.Args[2], os.Args[3])
		} else {
			generator.Init(os.Args[2], "")
		}
	case "model":
		if len(os.Args) == 2 {
			fmt.Println("Help Model Commands")
			return
		}
		generator.Model(os.Args[2])
	case "controller":
		if len(os.Args) == 2 {
			fmt.Println("Help Controller Commands")
			return
		}
		generator.Controller(os.Args[2])
	default:
		fmt.Println("Wrong command input!")
		fmt.Println("Help")
	}
}

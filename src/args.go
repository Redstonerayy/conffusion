package src

import (
	"fmt"
	"log"
	"os"
)

func ParseArguments(verbose bool) (string, map[string]string) {
	//parse processs arguments
	Argv := os.Args
	var Operation string
	var Flags = make(map[string]string)
	if len(Argv) > 1 {
		Operation = Argv[1]
	}
	for i := 2; i < len(Argv); i++ {
		if string(Argv[i][0]) == "-" {
			if string(Argv[i][1]) == "-" {
				//flag with argument, multiple letters, argument as string
				if len(Argv[i]) > 2 {
					Flags[Argv[i][2:]] = Argv[i+1]
					i++
				} else {
					log.Printf("Invalid long flag %s", Argv[i])
				}
			} else {
				//short flag, one letter, true or false
				if len(Argv[i]) == 2 {
					Flags[string(Argv[i][1])] = "true"
				} else {
					log.Printf("Invalid short flag %s", Argv[i])
				}
			}
		}
	}
	if verbose {
		fmt.Println(Operation)
		for val, key := range Flags {
			fmt.Printf("%s, %s\n", val, key)
		}
	}
	return Operation, Flags
}

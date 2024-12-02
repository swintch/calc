package main

import (
	"log"
	"os"

	"github.com/calc"
	"github.com/calc/CLIHandler"
)

func main() {
	handle := CLIHandler.NewCLIHandler(&calc.Addition{}, os.Stdout)
	err := handle.Handler(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}

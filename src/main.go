package main

import (
	"fmt"
	"os"

	"euclidean/trybasicrouter"
	"euclidean/trymux"
)

func main() {
	startupMessage := "Euclidean.GoSandbox Starting Up"
	fmt.Println(startupMessage)

	processParams()
}

func processParams() {
	args := os.Args
	if len(args) != 2 {
		printHelpMessage()
		return
	} else {
		param := args[1]
		switch param {
		case "basic":
			trybasicrouter.StartBasicServer()
		case "mux":
			trymux.StartMuxServer()
		default:
			printHelpMessage()
		}
	}
}

func printHelpMessage() {
	helpMessage := "\nEuclidean Go Sandbox is a playground for trying server functionality in GoLang\n"
	helpMessage += "Required input parameter options are \"basic\", \"mux\", and \"info\"\n"
	helpMessage += "\ti.e. `go run main.go basic`"
	fmt.Println(helpMessage)
}

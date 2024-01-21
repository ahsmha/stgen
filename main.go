package main

import (
	"fmt"
	"log"
	"os"
	"stgen/cmd"
)

func help() {

	help := `
    stgen - Copyright (C) 2024  ahsmha
	usage: stgen [args]

	args: 
		generate PATH				generate the html files at the given PATH
		serve [HOST:PORT]			serves 'site' folder
	`

	fmt.Println(help)
	return
}

func main() {

	args := os.Args

	hostPort := ":8282" // default host:port
	path := "./"        // default path for generating html files
	if len(args) == 3 {
		hostPort = args[2]
		path = args[2]
	} else if len(args) != 1 || len(args) != 2 {
		help()
		return
	}

	switch args[1] {
	case "generate":
		err := cmd.Generate(path)
		if err != nil {
			log.Fatalf("error while generating files: %+v \n", err)
		}

	case "serve":
		err := cmd.Serve(hostPort)
		if err != nil {
			log.Fatalf("error while serving files: %+v \n", err)
		}

	default:
		help()
	}

	return
}

package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var filePath string
	flag.StringVar(&filePath, "file", "", "Path to the .tef file")
	flag.Parse()

	if filePath == "" {
		log.Fatal("Please specify a .tef file using -file flag")
	}

	data, err := tabedit_parser.ParseTefFile(filePath)
	if err != nil {
		log.Fatalf("Error parsing TEF file: %v", err)
	}

	fmt.Println(data)
}

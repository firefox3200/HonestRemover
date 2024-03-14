package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flag for file path
	method := flag.String("method", "simple", "Method to use for rewriting the file. Options: simple, middle, high")
	filePath := flag.String("path", "", "Path to the file to rewrite")
	flag.Parse()

	// Check if file path is provided
	if *filePath == "" {
		fmt.Println("Please provide a file path using -path flag")
		return
	}

	switch *method {
	case "simple":
		remover := NewSimpleRemover(*filePath)
		remover.Remove()
	case "middle":
		remover := NewMiddleRemover(*filePath)
		remover.Remove()
	case "high":
		remover := NewHighRemover(*filePath)
		remover.Remove()
	default:
		fmt.Println("Invalid method provided. Please use one of the following: simple, middle, high")
		return
	}

	MakeEmpty(*filePath)
	newPath := RenameFileRandom(*filePath)
	DeleteFilePerm(newPath)
	fmt.Println("File rewritten with random data successfully.")
}

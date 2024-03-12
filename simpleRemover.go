package main

import (
	"crypto/rand"
	"log"
)

type simpleRemover struct {
	Name    string
	File    string
	Repeats int
}

func NewSimpleRemover(file string) Remover {
	return &simpleRemover{
		Name:    "simple",
		File:    file,
		Repeats: 3,
	}
}

func (r *simpleRemover) Remove() error {
	for i := 0; i < r.Repeats; i++ {
		file := GetFile(r.File)
		// Obtain file information
		fileInfo, err := file.Stat()
		if err != nil {
			log.Fatalf("Error getting file information: %v", err)
		}

		// Generate random data to rewrite the file
		dataSize := fileInfo.Size()
		randomData := make([]byte, dataSize)
		_, err = rand.Read(randomData)
		if err != nil {
			log.Fatalf("Error generating random data: %v", err)
		}

		// Write the random data to the file
		_, err = file.WriteAt(randomData, 0)
		if err != nil {
			log.Fatalf("Error writing random data to file: %v", err)
		}
		file.Close()
	}
	return nil
}

package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Remover interface {
	Remove() error
}

func GetFile(path string) (file *os.File) {
	// Open the file for writing
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	return file
}

func MakeEmpty(path string) {
	file := GetFile(path)
	file.Truncate(0)
	file.Close()
}

func DeleteFilePerm(path string) {
	err := os.Remove(path)
	if err != nil {
		log.Fatalf("Error deleting file: %v", err)
	}
}

func RenameFileRandom(path string) string {
	dirPath := filepath.Dir(path)
	newName := cryptoRandString(10)
	newPath := filepath.Join(dirPath, newName)
	os.Rename(path, newPath)
	return newPath
}

func cryptoRandString(i int) string {
	b := make([]byte, i)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatalf("Error generating random string: %v", err)
	}
	return fmt.Sprintf("%x", b)
}

var METHODS = [...]string{"simple", "middle", "high"}

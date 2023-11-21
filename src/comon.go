package main

import (
	"bufio"
	"log"
	"os"
)

func ReadLines(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Unable to open file %v", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

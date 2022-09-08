package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := os.Args[1]
	fmt.Println("Attempting to open file:", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}

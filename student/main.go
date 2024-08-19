package main

import (
	"fmt"
	"os"

	"guess-it-1/handler"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run main.go")
		return
	}

	handler.HandleInput()
}

package main

import (
	"log"
	"os"
)

func main() {
	content := []byte("Hello, Permissions!")
	err := os.WriteFile("example.txt", content, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

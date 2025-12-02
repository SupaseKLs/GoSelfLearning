package main

import (
	"os"
)

func main() {
	//Append context in file
	f, err := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.WriteString("\nAppeeded content"); err != nil {
		panic(err)
	}
	//Open a file
	// file, err := os.Open("example.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// log.Println("File opened succesfully")

	//Read context in file
	// content, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(content))

	//Add context in file
	// content := []byte("Hello, Go!")
	// err := os.WriteFile("example.txt", content, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

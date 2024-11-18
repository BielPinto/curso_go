package main

import (
	"bufio"
	"fmt"
	"os"
)

const nameFile = "file.txt"

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}
	sizeFiles, err := f.Write([]byte("Writing datas in the file"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("File created successfully! size: %d", sizeFiles)
	f.Close()
	//reading
	file, err := os.ReadFile(nameFile)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))

	//reading little by little to opening the file

	filetwo, err := os.Open(nameFile)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(filetwo)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))

	}

	// err = os.Remove(nameFile)
	if err != nil {
		panic(err)
	}

}

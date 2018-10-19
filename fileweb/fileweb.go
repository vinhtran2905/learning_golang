package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := " hello from go"

	file, err := os.Create("./fromString.txt")

	checkError(err)

	defer file.Close()

	ln, err := io.WriteString(file, content)
	checkError(err)

	bytes := []byte(content)

	ioutil.WriteFile("./writeByte.txt", bytes, 0644)

	fmt.Printf("All done with file of %v characters", ln)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

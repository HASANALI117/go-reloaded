package main

import (
	"fmt"
	"log"
	"os"
	"reboot"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Input and output files are needed")
		return
	} else if len(os.Args) > 3 {
		fmt.Println("Only input and output files are needed")
		return
	}

	file := os.Args[1]
	resultFile := os.Args[2]

	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	inputText := reboot.AllConversions(string(content))

	err = os.WriteFile(resultFile, []byte(inputText), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"jsontransformer/m/v2/jsontransformer"
	"os"
)

func main() {
	jsonFile, err := os.Open("input.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	fmt.Println("Successfully Opened input.json")

	inputJson, _ := os.ReadFile(jsonFile.Name())
	output, err := jsontransformer.Transform(inputJson)
	if err != nil {
		fmt.Println("Error returned during transformation", err)
	}
	fmt.Println("Final Output: ", output)
}

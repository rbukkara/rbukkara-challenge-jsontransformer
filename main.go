package main

import (
	"fmt"
	"jsontransformer/m/v2/fileio"
	"jsontransformer/m/v2/jsontransformer"
)

func main() {

	inputJson := fileio.ReadFile()

	output, err := jsontransformer.Transform(inputJson)
	if err != nil {
		fmt.Println("Error returned during transformation", err)
	}

	fileio.WriteFile(output)
}

package fileio

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

const (
	inputFileName  = "input.json"
	outputFileName = "output.json"
)

// ReadFile reads content from a file to be transformed
func ReadFile() io.Reader {

	jsonFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	fmt.Println("Successfully Opened file")

	inputJson, _ := os.ReadFile(jsonFile.Name())
	return bytes.NewReader(inputJson)

}

// WriteFile writes the json transformed content into a file
func WriteFile(reader io.Reader) {

	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	outputData := buf.String()

	f, err := os.OpenFile(outputFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println("error creating output file: ", err)
	}
	if _, err := f.Write(buf.Bytes()); err != nil {
		fmt.Println("Error writing output file: ", err)
	}
	fmt.Println("Final Output written to output file is: ", outputData)

}

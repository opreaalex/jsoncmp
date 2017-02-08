package main

import (
	"os"
	"errors"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func main() {
	fileNames, err := getInputFileNames()
	if err != nil {
		panic(err)
	}

	for _, fileName := range fileNames {
		fileBytes, err := ioutil.ReadFile(fileName)
		if err != nil {
			panic(err)
		}

		var jsonMap map[string]interface{}
		err = json.Unmarshal(fileBytes, &jsonMap)
		if err != nil {
			panic(err)
		}

		fmt.Println(jsonMap)
	}
}

func getInputFileNames() ([]string, error) {
	// Exclude the runnable argument
	args := os.Args[1:]
	if len(args) < 2 {
		return nil, errors.New("You must provide at least two files")
	}
	return args, nil
}

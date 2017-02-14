package main

import (
	"os"
	"errors"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

type jsonInfo struct {
	fileName string
	jsonMap map[string]interface{}
}

func main() {
	fileNames, err := getInputFileNames()
	if err != nil {
		panic(err)
	}

	// Read the JSON from the file into maps
	fstJsonMap := getJsonMap(fileNames[0])
	secJsonMap := getJsonMap(fileNames[1])

	// Get get root keys from the JSON maps
	fstJsonKeys := getKeys(fstJsonMap)
	secJsonKeys := getKeys(secJsonMap)

	// Find out the difference between the keys
	fstDiffs := getDifferentBetween(fstJsonKeys, secJsonKeys)
	secDiffs := getDifferentBetween(secJsonKeys, fstJsonKeys)

	// Print the results
	fmt.Print(fmt.Sprintf("File %s: ", fileNames[0]))
	fmt.Println(fstDiffs)

	fmt.Print(fmt.Sprintf("File %s: ", fileNames[1]))
	fmt.Println(secDiffs)
}

func getInputFileNames() ([]string, error) {
	// Exclude the runnable argument
	args := os.Args[1:]
	if len(args) != 2 {
		return nil, errors.New("You must provide two files")
	}
	return args, nil
}

func getJsonMap(fileName string) map[string]interface{} {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	var jsonMap map[string]interface{}
	err = json.Unmarshal(fileBytes, &jsonMap)
	if err != nil {
		panic(err)
	}
	return jsonMap
}

func getKeys(jsonMap map[string]interface{}) []string {
	keys := make([]string, 0, len(jsonMap))
	for jsonKey, _ := range jsonMap {
		keys = append(keys, jsonKey)
	}
	return keys
}

func getDifferentBetween(first []string, second []string) []string {
	diff := make([]string, 0, len(first))
	for _, a := range first {
		found := false
		for _, b := range second {
			if a == b {
				found = true
				break
			}
		}
		if !found {
			diff = append(diff, a)
		}
	}
	return diff
}

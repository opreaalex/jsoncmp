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
	aFileName, bFileName := getInputFileNames()

	// Read the JSON from the file into maps
	aJsonMap := getJsonMap(aFileName)
	bJsonMap := getJsonMap(bFileName)

	// Get the diff maps
	// Compare the first json with the second
	// And compare the second with the first
	aDiffMap := getDiffMap(aJsonMap, bJsonMap)
	bDiffMap := getDiffMap(bJsonMap, aJsonMap)

	// Construct the result JSON
	resultJsonMap := make(map[string]interface{})
	resultJsonMap[aFileName] = aDiffMap
	resultJsonMap[bFileName] = bDiffMap
	resultJson := getJsonString(resultJsonMap)

	// Print the result
	fmt.Println(resultJson)
}

func getInputFileNames() (string, string) {
	args := os.Args[1:] // Exclude de runnable argument
	if len(args) != 2 {
		panic(errors.New("You must provide two files"))
	}
	return args[0], args[1]
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

func getDiffMap(aMap map[string]interface{}, bMap map[string]interface{}) map[string]interface{} {
	diff := make(map[string]interface{})
	for aKey, aVal := range aMap {
		found := false
		for bKey, bVal := range bMap {
			if aKey == bKey {
				found = true
				aInnerMap, aIsInnerMap := aVal.(map[string]interface{})
				bInnerMap, bIsInnerMap := bVal.(map[string]interface{})
				if aIsInnerMap && bIsInnerMap {
					innerDiff := getDiffMap(aInnerMap, bInnerMap)
					if len(innerDiff) > 0 {
						diff[aKey] = innerDiff
					}
				}
				break
			}
		}
		if !found {
			diff[aKey] = aVal
		}
	}
	return diff
}

func getJsonString(jsonMap map[string]interface{}) string {
	jsonBytes, err := json.MarshalIndent(jsonMap, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}

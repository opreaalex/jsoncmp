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

	jsonInfoArr := make([]jsonInfo, 0, len(fileNames))
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

		ji := jsonInfo{
			fileName: fileName,
			jsonMap: jsonMap,
		}

		jsonInfoArr = append(jsonInfoArr, ji)
	}

	keys := getKeys(jsonInfoArr)
	for k, v := range keys {
		fmt.Println(fmt.Sprint("For the file %s, we jave the keys: ", k))
		for _, field := range v {
			fmt.Println(field)
		}
		fmt.Println("----------")
	}
}

func getInputFileNames() ([]string, error) {
	// Exclude the runnable argument
	args := os.Args[1:]
	if len(args) != 2 {
		return nil, errors.New("You must provide two files")
	}
	return args, nil
}

func getKeys(jsonInfoArr []jsonInfo) map[string][]string {
	keyMap := make(map[string][]string)
	for _, ji := range jsonInfoArr {
		length := len(ji.jsonMap)
		keys := make([]string, 0, length)
		for key, _ := range ji.jsonMap {
			keys = append(keys, key)
		}
		keyMap[ji.fileName] = keys
	}
	return keyMap
}

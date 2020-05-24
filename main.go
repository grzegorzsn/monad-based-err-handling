package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func main() {
	err := runEverything()
	if err != nil {
		fmt.Println("Runtime error:", err)
		return
	}
	fmt.Println("Everything fine")
}

func runEverything() error {
	fileData, err := fetchDataFromFile()
	if err != nil {
		return err
	}
	dbData, err := fetchDataFromDb()
	if err != nil {
		return err
	}
	processedData := processData(fileData, dbData) // no error as we expect input is already perfectly validated
	err = saveProcessedData(processedData)
	if err != nil {
		return err
	}
	fmt.Println("Processing done")
	return nil
}

func fetchDataFromFile() (interface{}, error) {
	bytes, err := ioutil.ReadFile("some_filename")
	if err != nil {
		return nil, err
	}
	data, err := parseData(bytes)
	if err != nil {
		return nil, err
	}
	err = validate(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// -----------------------------------------------
// Code below is not implemented. You can skip it.
// -----------------------------------------------

func validate(data interface{}) error {
	return errors.New("not implemented")
}

func parseData(bytes []byte) (interface{}, error) {
	return nil, errors.New("not implemented")
}

func fetchDataFromDb() (interface{}, error) {
	return nil, errors.New("not implemented")
}

func saveProcessedData(data interface{}) error {
	return errors.New("not implemented")
}

func processData(data interface{}, dbData interface{}) interface{} {
	return data
}

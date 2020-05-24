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
	err := runEverythingCompose(
		fetchData,
		processData,
		saveProcessedData,
	)()
	return err
}

func fetchData() (interface{}, interface{}, error) {
	fileData, err := fetchDataFromFile()
	if err != nil {
		return nil, nil, err
	}
	dbData, err := fetchDataFromDb()
	if err != nil {
		return nil, nil, err
	}
	return fileData, dbData, nil
}

func fetchDataFromFile() (interface{}, error) {
	return fetchDataFromFileCompose(
		ioutil.ReadFile,
		parseData,
		validate,
	)("some_filename")
}

// -----------------------------------------------
// Code below is not implemented. You can skip it.
// -----------------------------------------------

func validate(data interface{}) (interface{}, error) {
	return nil, errors.New("not implemented")
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

func processData(data interface{}, dbData interface{}) (interface{}, error) {
	return data, nil
}

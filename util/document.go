package util

import (
	"fmt"
	"os"
)

func GetJSONFile(fileName string) (*os.File, error) {
	jsonFile, err := os.Open(fileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	return jsonFile, nil
}

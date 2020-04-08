package common

import (
	"os"

	"io/ioutil"
)

// Convert file to string
func ConvertFileToStr(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	buffer, err := ioutil.ReadAll(file)
	return string(buffer), nil
}

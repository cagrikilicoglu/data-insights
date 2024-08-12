package file

import (
	"data-insights/kit/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// ProcessJSONFile reads and parses the JSON file
func GetRawDataFromFile(filePath string) ([]model.Insight, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// todo change readall
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file content: %v", err)
	}

	//log.Printf("Content of file %s: %s\n", filePath, string(content))

	var data []model.Insight
	if err := json.Unmarshal(content, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}
	//fmt.Printf("Parsed data from file %s: %+v\n", filePath, data)
	return data, nil
}

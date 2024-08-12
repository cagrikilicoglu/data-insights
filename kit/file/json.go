package file

import (
	"data-insights/kit/common"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// GetRawDataFromFile reads and parses the JSON file at the given filePath.
// It returns a slice of Insight objects or an error if something goes wrong.
func GetRawDataFromFile(filePath string) ([]common.Insight, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Read the file's content
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file content: %v", err)
	}

	var data []common.Insight
	if err := json.Unmarshal(content, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return data, nil
}

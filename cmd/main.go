package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	fileDirectory := os.Getenv("file_dir")
	if fileDirectory == "" {
		log.Fatal("There is no file path supplied")
	}

	files, err := FilePathWalkDir(fileDirectory)
	if err != nil {
		panic(err)
	}

	// Open the CSV file
	for _, fileSource := range files {
		err := processJSONFile(fileSource)
		if err != nil {
			log.Fatal("can not unmarshall json data")
		}
	}
}

func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// processJSONFile reads and parses the JSON file
func processJSONFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read file content: %v", err)
	}

	log.Printf("Content of file %s: %s\n", filePath, string(content))

	var data []Insight
	if err := json.Unmarshal(content, &data); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}
	fmt.Printf("Parsed data from file %s: %+v\n", filePath, data)
	return nil
}

type Insight struct {
	Country                string `json:"Country"`
	DeviceCategory         string `json:"DeviceCategory"`
	EngagementRate         string `json:"EngagementRate"`
	LandingPage            string `json:"LandingPage"`
	NewUsers               int    `json:"NewUsers"`
	ScreenPageViews        int    `json:"ScreenPageViews"`
	SessionMedium          string `json:"SessionMedium"`
	Sessions               int    `json:"Sessions"`
	TotalUsers             int    `json:"TotalUsers"`
	UserEngagementDuration int    `json:"UserEngagementDuration"`
	Date                   string `json:"date"`
}

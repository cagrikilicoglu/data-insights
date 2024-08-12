package main

import (
	"data-insights/kit/ai"
	"data-insights/kit/file"
	"data-insights/kit/metrics"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const thresholdDataPointNumber int = 100

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	fileDirectory := os.Getenv("FILE_DIR")
	if fileDirectory == "" {
		log.Fatal("There is no file path supplied")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable is required")
	}

	files, err := file.FilePathWalkDir(fileDirectory)
	if err != nil {
		panic(err)
	}

	for _, fileSource := range files {
		data, err := file.ProcessJSONFile(fileSource)
		if err != nil {
			log.Printf("Error processing file %s: %v\n", fileSource, err)
			continue
		}
		importantMetrics := metrics.GetImportantMetrics(data)
		prompt := ai.CreatePrompt(importantMetrics)
		fmt.Println(prompt)
		//overallMetrics := metrics.CalculateOverallMetrics(data)
		//overallMetricsJSON, err := json.MarshalIndent(overallMetrics, "", "  ")
		//if err != nil {
		//	fmt.Println("Error marshalling to JSON:", err)
		//	return
		//}
		//fmt.Println(string(overallMetricsJSON))
		//
		//averageEngagementRatesByCountry := metrics.AggregateMetricsByBreakdown(data, metrics.COUNTRY, thresholdDataPointNumber)
		//top5AverageEngagementRatesByCountryJSON, err := json.MarshalIndent(metrics.GetTopElements(averageEngagementRatesByCountry, 5), "", "  ")
		//if err != nil {
		//	fmt.Println("Error marshalling to JSON:", err)
		//	return
		//}
		//fmt.Println(string(top5AverageEngagementRatesByCountryJSON))
		//bottom5AverageEngagementRatesByCountryJSON, err := json.MarshalIndent(metrics.GetBottomElements(averageEngagementRatesByCountry, 5), "", "  ")
		//if err != nil {
		//	fmt.Println("Error marshalling to JSON:", err)
		//	return
		//}
		//fmt.Println(string(bottom5AverageEngagementRatesByCountryJSON))
		//
		//averageDataForDevices := metrics.AggregateMetricsByBreakdown(data, metrics.DEVICE, thresholdDataPointNumber)
		//top5AverageDataForDevicesJSON, err := json.MarshalIndent(metrics.GetTopElements(averageDataForDevices, 3), "", "  ")
		//if err != nil {
		//	fmt.Println("Error marshalling to JSON:", err)
		//	return
		//}
		//fmt.Println(string(top5AverageDataForDevicesJSON))
		//bottom5AverageDataForDevicesJSON, err := json.MarshalIndent(metrics.GetBottomElements(averageDataForDevices, 3), "", "  ")
		//if err != nil {
		//	fmt.Println("Error marshalling to JSON:", err)
		//	return
		//}
		//fmt.Println(string(bottom5AverageDataForDevicesJSON))
		//
		//averageDataForPages := metrics.AggregateMetricsByBreakdown(data, metrics.PAGE, thresholdDataPointNumber)
		//top5AverageDataForPagesJSON, err := json.MarshalIndent(metrics.GetTopElements(averageDataForPages, 5), "", "  ")
		//if err != nil {
		//	fmt.Println("Error marshalling to JSON:", err)
		//	return
		//}
		//fmt.Println(string(top5AverageDataForPagesJSON))
		//bottom5AverageDataForPagesJSON, err := json.MarshalIndent(metrics.GetBottomElements(averageDataForPages, 5), "", "  ")
		//if err != nil {
		//	fmt.Println("Error marshalling to JSON:", err)
		//	return
		//}
		//fmt.Println(string(bottom5AverageDataForPagesJSON))
		//
		//averageDataForSessions := metrics.AggregateMetricsByBreakdown(data, metrics.MEDIUM, thresholdDataPointNumber)
		//top5AverageDataForSessionsJSON, err := json.MarshalIndent(metrics.GetTopElements(averageDataForSessions, 2), "", "  ")
		//if err != nil {
		//	fmt.Println("Error marshalling to JSON:", err)
		//	return
		//}
		//fmt.Println(string(top5AverageDataForSessionsJSON))
		//bottom5AverageDataForSessionsJSON, err := json.MarshalIndent(metrics.GetBottomElements(averageDataForSessions, 2), "", "  ")
		//if err != nil {
		//	fmt.Println("Error marshalling to JSON:", err)
		//	return
		//}
		//fmt.Println(string(bottom5AverageDataForSessionsJSON))

		insights, err := ai.GetInsightsFromLLM(apiKey, importantMetrics)
		if err != nil {
			log.Printf("Error getting insights from LLM: %v\n", err)
			continue
		}

		//insightsMarsh, err := json.Marshal(insights)
		//if err != nil {
		//	log.Printf("error when marshalling")
		//	return
		//}
		var userMetricsWithInsights ai.UserMetricsWithInsights
		err = json.Unmarshal([]byte(insights), &userMetricsWithInsights)
		if err != nil {
			log.Printf("error when unmarshalling")
			return
		}

		fmt.Printf("Insights for file %s: %s\n", fileSource, insights)
		fmt.Println("User Insights for file %s: %s\n", fileSource, userMetricsWithInsights)
	}
}

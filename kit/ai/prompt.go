package ai

import (
	"data-insights/kit/model"
	"fmt"
)

//type AggregatedMetricsWithInsight struct {
//	model.AggregatedMetricsList `json:"aggregated_metrics"`
//	AIInsight                   string `json:"ai_insight"`
//}
//
//type OverallMetricsWithInsight struct {
//	model.OverallMetrics `json:"overall_metrics"`
//	AIInsight            string `json:"ai_insight"`
//}
//
//type UserMetricsWithInsights struct {
//	OverallMetrics                         OverallMetricsWithInsight    `json:"overall_metrics"`
//	Top5CountriesWithHighestEngagementRate AggregatedMetricsWithInsight `json:"top_5_countries_with_highest_engagement_rate"`
//	Top5CountriesWithLowestEngagementRate  AggregatedMetricsWithInsight `json:"top_5_countries_with_lowest_engagement_rate"`
//	BounceRatesByDevices                   AggregatedMetricsWithInsight `json:"bounce_rates_by_devices"`
//	Top5PagesWithHighestNoOfSessions       AggregatedMetricsWithInsight `json:"top_5_pages_with_highest_no_of_sessions"`
//	Top5PagesWithLowestNoOfSessions        AggregatedMetricsWithInsight `json:"top_5_pages_with_lowest_no_of_sessions"`
//	AverageSessionDurationsByDevices       AggregatedMetricsWithInsight `json:"average_session_durations_by_devices"`
//}

type OverallMetricsWithInsight struct {
	OverallEngagementRate  string `json:"overall_engagement_rate"`
	AverageSessionDuration string `json:"average_session_duration"`
	BounceRate             string `json:"bounce_rate"`
	PagesPerSession        string `json:"pages_per_session"`
	NewUserPercentage      string `json:"new_user_percentage"`
	SessionPerUser         string `json:"session_per_user"`
	AIInsight              string `json:"ai_insight"`
}

type AggregatedMetric struct {
	Name                   string `json:"name"`
	AverageEngagementRate  string `json:"average_engagement_rate,omitempty"`
	BounceRate             string `json:"bounce_rate,omitempty"`
	TotalSessions          string `json:"total_sessions,omitempty"`
	AverageSessionDuration string `json:"average_session_duration,omitempty"`
}

type AggregatedMetricsWithInsight struct {
	AIInsight         string             `json:"ai_insight"`
	AggregatedMetrics []AggregatedMetric `json:"aggregated_metrics"`
}

type UserMetricsWithInsights struct {
	OverallMetrics                         OverallMetricsWithInsight    `json:"overall_metrics"`
	Top5CountriesWithHighestEngagementRate AggregatedMetricsWithInsight `json:"top_5_countries_with_highest_engagement_rate"`
	Top5CountriesWithLowestEngagementRate  AggregatedMetricsWithInsight `json:"top_5_countries_with_lowest_engagement_rate"`
	BounceRatesByDevices                   AggregatedMetricsWithInsight `json:"bounce_rates_by_devices"`
	Top5PagesWithHighestNoOfSessions       AggregatedMetricsWithInsight `json:"top_5_pages_with_highest_no_of_sessions"`
	Top5PagesWithLowestNoOfSessions        AggregatedMetricsWithInsight `json:"top_5_pages_with_lowest_no_of_sessions"`
	AverageSessionDurationsByDevices       AggregatedMetricsWithInsight `json:"average_session_durations_by_devices"`
}

func CreatePrompt(metrics model.UserMetrics) string {
	prompt := "Analyze the following metrics and provide insights for each group of metrics as a whole. The insights should be included in the 'ai_insight' field for each metric group.\n\n"

	// Add overall metrics
	prompt += "Overall Metrics:\n"
	prompt += fmt.Sprintf("  - Overall Engagement Rate: %.2f%%\n", metrics.OverallMetrics.OverallEngagementRate*100)
	prompt += fmt.Sprintf("  - Average Session Duration: %.2f seconds\n", metrics.OverallMetrics.AverageSessionDuration)
	prompt += fmt.Sprintf("  - Bounce Rate: %.2f%%\n", metrics.OverallMetrics.BounceRate)
	prompt += fmt.Sprintf("  - Pages Per Session: %.2f\n", metrics.OverallMetrics.PagesPerSession)
	prompt += fmt.Sprintf("  - New User Percentage: %.2f%%\n", metrics.OverallMetrics.NewUserPercentage)
	prompt += fmt.Sprintf("  - Session Per User: %.2f\n", metrics.OverallMetrics.SessionPerUser)

	// Add top 5 countries with highest engagement rate
	prompt += "\nTop 5 Countries with Highest Engagement Rate:\n"
	for _, metric := range metrics.Top5CountriesWithHighestEngagementRate {
		prompt += fmt.Sprintf("  - %s: %.2f%% engagement rate\n", metric.Name, metric.AverageEngagementRate*100)
	}

	// Add top 5 countries with lowest engagement rate
	prompt += "\nTop 5 Countries with Lowest Engagement Rate:\n"
	for _, metric := range metrics.Top5CountriesWithLowestEngagementRate {
		prompt += fmt.Sprintf("  - %s: %.2f%% engagement rate\n", metric.Name, metric.AverageEngagementRate*100)
	}

	// Add bounce rates by devices
	prompt += "\nBounce Rates by Devices:\n"
	for _, metric := range metrics.BounceRatesByDevices {
		prompt += fmt.Sprintf("  - %s: %.2f%% bounce rate\n", metric.Name, metric.BounceRate)
	}

	// Add top 5 pages with highest number of sessions
	prompt += "\nTop 5 Pages with Highest Number of Sessions:\n"
	for _, metric := range metrics.Top5PagesWithHighestNoOfSessions {
		prompt += fmt.Sprintf("  - %s: %d sessions\n", metric.Name, metric.TotalSessions)
	}

	// Add top 5 pages with lowest number of sessions
	prompt += "\nTop 5 Pages with Lowest Number of Sessions:\n"
	for _, metric := range metrics.Top5PagesWithLowestNoOfSessions {
		prompt += fmt.Sprintf("  - %s: %d sessions\n", metric.Name, metric.TotalSessions)
	}

	// Add average session durations by devices
	prompt += "\nAverage Session Durations by Devices:\n"
	for _, metric := range metrics.AverageSessionDurationsByDevices {
		prompt += fmt.Sprintf("  - %s: %.2f seconds average session duration\n", metric.Name, metric.AverageSessionDuration)
	}

	prompt += "\nPlease provide insights for each group of metrics as a whole in the 'ai_insight' field. The output should be in the following JSON structure without any additional words:\n"
	prompt += "{\n"
	prompt += "  \"overall_metrics\": {\n"
	prompt += "    \"overall_engagement_rate\": \"value\",\n"
	prompt += "    \"average_session_duration\": \"value\",\n"
	prompt += "    \"bounce_rate\": \"value\",\n"
	prompt += "    \"pages_per_session\": \"value\",\n"
	prompt += "    \"new_user_percentage\": \"value\",\n"
	prompt += "    \"session_per_user\": \"value\",\n"
	prompt += "    \"ai_insight\": \"insight\"\n"
	prompt += "  },\n"
	prompt += "  \"top_5_countries_with_highest_engagement_rate\": {\n"
	prompt += "    \"ai_insight\": \"insight\",\n"
	prompt += "    \"aggregated_metrics\": [\n"
	prompt += "      {\n"
	prompt += "        \"name\": \"country_name\",\n"
	prompt += "        \"average_engagement_rate\": \"value\"\n"
	prompt += "      }\n"
	prompt += "    ]\n"
	prompt += "  },\n"
	prompt += "  \"top_5_countries_with_lowest_engagement_rate\": {\n"
	prompt += "    \"ai_insight\": \"insight\",\n"
	prompt += "    \"aggregated_metrics\": [\n"
	prompt += "      {\n"
	prompt += "        \"name\": \"country_name\",\n"
	prompt += "        \"average_engagement_rate\": \"value\"\n"
	prompt += "      }\n"
	prompt += "    ]\n"
	prompt += "  },\n"
	prompt += "  \"bounce_rates_by_devices\": {\n"
	prompt += "    \"ai_insight\": \"insight\",\n"
	prompt += "    \"aggregated_metrics\": [\n"
	prompt += "      {\n"
	prompt += "        \"name\": \"device_category\",\n"
	prompt += "        \"bounce_rate\": \"value\"\n"
	prompt += "      }\n"
	prompt += "    ]\n"
	prompt += "  },\n"
	prompt += "  \"top_5_pages_with_highest_no_of_sessions\": {\n"
	prompt += "    \"ai_insight\": \"insight\",\n"
	prompt += "    \"aggregated_metrics\": [\n"
	prompt += "      {\n"
	prompt += "        \"name\": \"page_name\",\n"
	prompt += "        \"total_sessions\": \"value\"\n"
	prompt += "      }\n"
	prompt += "    ]\n"
	prompt += "  },\n"
	prompt += "  \"top_5_pages_with_lowest_no_of_sessions\": {\n"
	prompt += "    \"ai_insight\": \"insight\",\n"
	prompt += "    \"aggregated_metrics\": [\n"
	prompt += "      {\n"
	prompt += "        \"name\": \"page_name\",\n"
	prompt += "        \"total_sessions\": \"value\"\n"
	prompt += "      }\n"
	prompt += "    ]\n"
	prompt += "  },\n"
	prompt += "  \"average_session_durations_by_devices\": {\n"
	prompt += "    \"ai_insight\": \"insight\",\n"
	prompt += "    \"aggregated_metrics\": [\n"
	prompt += "      {\n"
	prompt += "        \"name\": \"device_category\",\n"
	prompt += "        \"average_session_duration\": \"value\"\n"
	prompt += "      }\n"
	prompt += "    ]\n"
	prompt += "  }\n"
	prompt += "}\n"

	return prompt
}

//
//func createPrompt(data []model.Insight) string {
//	prompt := "Analyze the following data and provide insights:\n\n"
//	for i, d := range data {
//		prompt += fmt.Sprintf("Country: %s, DeviceCategory: %s, EngagementRate: %s, LandingPage: %s, NewUsers: %d, ScreenPageViews: %d, SessionMedium: %s, Sessions: %d, TotalUsers: %d, UserEngagementDuration: %d, Date: %s\n",
//			d.Country, d.DeviceCategory, d.EngagementRate, d.LandingPage, d.NewUsers, d.ScreenPageViews, d.SessionMedium, d.Sessions, d.TotalUsers, d.UserEngagementDuration, d.Date)
//		if i == 100 {
//			break
//		}
//	}
//	//fmt.Println(prompt)
//	//for i:= 0 ; i<100 ; i++ {
//	//		prompt += fmt.Sprintf("Country: %s, DeviceCategory: %s, EngagementRate: %s, LandingPage: %s, NewUsers: %d, ScreenPageViews: %d, SessionMedium: %s, Sessions: %d, TotalUsers: %d, UserEngagementDuration: %d, Date: %s\n",
//	//			d.Country, d.DeviceCategory, d.EngagementRate, d.LandingPage, d.NewUsers, d.ScreenPageViews, d.SessionMedium, d.Sessions, d.TotalUsers, d.UserEngagementDuration, d.Date)
//	//	}
//	//}
//	return prompt
//}

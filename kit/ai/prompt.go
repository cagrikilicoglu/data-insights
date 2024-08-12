package ai

import (
	"data-insights/kit/common"
	"fmt"
)

// createPrompt generates a formatted string prompt based on the provided UserMetrics.
// It uses a template to organize metrics into different sections and returns a complete
// prompt string that is ready to be sent to the AI model for insights.
func createPrompt(metrics common.UserMetrics) string {

	top5Highest := formatEngagementRateList(metrics.Top5CountriesWithHighestEngagementRate)
	top5Lowest := formatEngagementRateList(metrics.Top5CountriesWithLowestEngagementRate)
	bounceRates := formatBounceRateList(metrics.BounceRatesByDevices)
	top5HighestSessions := formatSessionCountList(metrics.Top5PagesWithHighestNoOfSessions)
	top5LowestSessions := formatSessionCountList(metrics.Top5PagesWithLowestNoOfSessions)
	avgSessionDurations := formatSessionDurationList(metrics.AverageSessionDurationsByDevices)

	return fmt.Sprintf(PromptFormat,
		metrics.OverallMetrics.OverallEngagementRate*100,
		metrics.OverallMetrics.AverageSessionDuration,
		metrics.OverallMetrics.BounceRate,
		metrics.OverallMetrics.PagesPerSession,
		metrics.OverallMetrics.NewUserPercentage,
		metrics.OverallMetrics.SessionPerUser,
		top5Highest,
		top5Lowest,
		bounceRates,
		top5HighestSessions,
		top5LowestSessions,
		avgSessionDurations,
	)
}

func formatEngagementRateList(metrics []common.AggregatedMetrics) string {
	if len(metrics) == 0 {
		return "  - No data available\n"
	}
	var result string
	for _, metric := range metrics {
		result += fmt.Sprintf("  - %s: %.2f%% engagement rate\n", metric.Name, metric.AverageEngagementRate*100)
	}
	return result
}

func formatBounceRateList(metrics []common.AggregatedMetrics) string {
	if len(metrics) == 0 {
		return "  - No data available\n"
	}
	var result string
	for _, metric := range metrics {
		result += fmt.Sprintf("  - %s: %.2f%% bounce rate\n", metric.Name, metric.BounceRate)
	}
	return result
}

func formatSessionCountList(metrics []common.AggregatedMetrics) string {
	if len(metrics) == 0 {
		return "  - No data available\n"
	}
	var result string
	for _, metric := range metrics {
		result += fmt.Sprintf("  - %s: %d sessions\n", metric.Name, metric.TotalSessions)
	}
	return result
}

func formatSessionDurationList(metrics []common.AggregatedMetrics) string {
	if len(metrics) == 0 {
		return "  - No data available\n"
	}
	var result string
	for _, metric := range metrics {
		result += fmt.Sprintf("  - %s: %.2f seconds average session duration\n", metric.Name, metric.AverageSessionDuration)
	}
	return result
}

const PromptFormat = `
Analyze the following metrics and provide insights for each group of metrics as a whole. The insights should be included in the 'ai_insight' field for each metric group.

Overall Metrics:
  - Overall Engagement Rate: %.2f%%
  - Average Session Duration: %.2f seconds
  - Bounce Rate: %.2f%%
  - Pages Per Session: %.2f
  - New User Percentage: %.2f%%
  - Session Per User: %.2f

Top 5 Countries with Highest Engagement Rate:
%s

Top 5 Countries with Lowest Engagement Rate:
%s

Bounce Rates by Devices:
%s

Top 5 Pages with Highest Number of Sessions:
%s

Top 5 Pages with Lowest Number of Sessions:
%s

Average Session Durations by Devices:
%s

Please provide insights for each group of metrics as a whole in the 'ai_insight' field. The output should be in the following JSON structure without any additional words:
{
  "overall_metrics": {
    "overall_engagement_rate": "value",
    "average_session_duration": "value",
    "bounce_rate": "value",
    "pages_per_session": "value",
    "new_user_percentage": "value",
    "session_per_user": "value",
    "ai_insight": "insight"
  },
  "top_5_countries_with_highest_engagement_rate": {
    "ai_insight": "insight",
    "aggregated_metrics": [
      {
        "name": "country_name",
        "average_engagement_rate": "value"
      }
    ]
  },
  "top_5_countries_with_lowest_engagement_rate": {
    "ai_insight": "insight",
    "aggregated_metrics": [
      {
        "name": "country_name",
        "average_engagement_rate": "value"
      }
    ]
  },
  "bounce_rates_by_devices": {
    "ai_insight": "insight",
    "aggregated_metrics": [
      {
        "name": "device_category",
        "bounce_rate": "value"
      }
    ]
  },
  "top_5_pages_with_highest_no_of_sessions": {
    "ai_insight": "insight",
    "aggregated_metrics": [
      {
        "name": "page_name",
        "total_sessions": "value"
      }
    ]
  },
  "top_5_pages_with_lowest_no_of_sessions": {
    "ai_insight": "insight",
    "aggregated_metrics": [
      {
        "name": "page_name",
        "total_sessions": "value"
      }
    ]
  },
  "average_session_durations_by_devices": {
    "ai_insight": "insight",
    "aggregated_metrics": [
      {
        "name": "device_category",
        "average_session_duration": "value"
      }
    ]
  }
}
`

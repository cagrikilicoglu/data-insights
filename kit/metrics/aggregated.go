package metrics

import (
	"data-insights/kit/model"
)

func AggregateMetricsByBreakdown(data []model.Insight, breakdown model.Breakdown, threshold int) model.AggregatedMetricsList {
	metricsMap := make(map[string][]model.Insight)

	// Aggregate data based on the breakdown (Country, DeviceCategory, LandingPage, SessionMedium)
	for _, insight := range data {
		var key string
		switch breakdown {
		case model.COUNTRY:
			key = insight.Country
		case model.DEVICE:
			key = insight.DeviceCategory
		case model.PAGE:
			key = insight.LandingPage
		case model.MEDIUM:
			key = insight.SessionMedium
		}
		metricsMap[key] = append(metricsMap[key], insight)
	}

	// Calculate metrics for each breakdown category that meets the threshold
	var aggregatedMetrics []model.AggregatedMetrics
	for name, insights := range metricsMap {
		if len(insights) >= threshold && name != model.NOTSET {
			var totalEngagementRate, totalSessionDuration, totalBounceRate float64
			var totalSessions, totalPageViews, totalNewUsers, totalUsers int

			for _, insight := range insights {
				totalEngagementRate += parseStringToFloat(insight.EngagementRate)
				totalSessionDuration += float64(insight.UserEngagementDuration)
				if insight.Sessions > 0 && insight.ScreenPageViews == 1 {
					totalBounceRate += 1
				}
				totalPageViews += insight.ScreenPageViews
				totalSessions += insight.Sessions
				totalNewUsers += insight.NewUsers
				totalUsers += insight.TotalUsers
			}

			averageEngagementRate := totalEngagementRate / float64(len(insights))
			averageSessionDuration := totalSessionDuration / float64(totalSessions)
			bounceRate := (totalBounceRate / float64(totalSessions)) * 100
			averageEngagementDuration := totalSessionDuration / float64(totalPageViews)

			aggregatedMetrics = append(aggregatedMetrics, model.AggregatedMetrics{
				Name:                      name,
				AverageEngagementRate:     averageEngagementRate,
				TotalSessions:             totalSessions,
				TotalPageViews:            totalPageViews,
				AverageSessionDuration:    averageSessionDuration,
				BounceRate:                bounceRate,
				TotalNewUsers:             totalNewUsers,
				TotalUsers:                totalUsers,
				AverageEngagementDuration: averageEngagementDuration,
				DataPointCount:            len(insights),
			})
		}
	}
	return aggregatedMetrics
}

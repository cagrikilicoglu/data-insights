package metrics

import "data-insights/kit/common"

const singlePageView = 1

// AggregateMetricsByBreakdown aggregates metrics based on the given breakdown (e.g., Country, DeviceCategory, etc.).
// It applies a threshold to include only those categories with sufficient data points and returns a list of aggregated metrics.
func AggregateMetricsByBreakdown(data []common.Insight, breakdown common.Breakdown, threshold int) common.AggregatedMetricsList {
	metricsMap := make(map[string][]common.Insight)

	// Group data based on the breakdown (Country, DeviceCategory, LandingPage, SessionMedium)
	for _, insight := range data {
		var key string
		switch breakdown {
		case common.COUNTRY:
			key = insight.Country
		case common.DEVICE:
			key = insight.DeviceCategory
		case common.PAGE:
			key = insight.LandingPage
		case common.MEDIUM:
			key = insight.SessionMedium
		}
		metricsMap[key] = append(metricsMap[key], insight)
	}

	// Calculate aggregated metrics for each group that meets the threshold
	var aggregatedMetrics []common.AggregatedMetrics
	for name, insights := range metricsMap {
		if len(insights) >= threshold && name != common.NOTSET {
			var totalEngagementRate, totalSessionDuration, totalBounceRate float64
			var totalSessions, totalPageViews, totalNewUsers, totalUsers int

			// Sum up the metrics for the current group
			for _, insight := range insights {
				totalEngagementRate += parseStringToFloat(insight.EngagementRate)
				totalSessionDuration += float64(insight.UserEngagementDuration)
				if insight.Sessions > 0 && insight.ScreenPageViews == singlePageView {
					totalBounceRate += 1
				}
				totalPageViews += insight.ScreenPageViews
				totalSessions += insight.Sessions
				totalNewUsers += insight.NewUsers
				totalUsers += insight.TotalUsers
			}

			// Calculate averages and derived metrics
			averageEngagementRate := totalEngagementRate / float64(len(insights))
			averageSessionDuration := totalSessionDuration / float64(totalSessions)
			bounceRate := (totalBounceRate / float64(totalSessions)) * 100

			var averageEngagementDuration float64
			if totalPageViews > 0 {
				averageEngagementDuration = totalSessionDuration / float64(totalPageViews)
			}

			// Append the aggregated metrics for the current group
			aggregatedMetrics = append(aggregatedMetrics, common.AggregatedMetrics{
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

package metrics

import (
	"data-insights/kit/model"
	"sort"
)

const NOTSET string = "(not set)"

type Breakdown string

const (
	COUNTRY Breakdown = "Country"
	DEVICE  Breakdown = "DeviceCategory"
	PAGE    Breakdown = "LandingPage"
	MEDIUM  Breakdown = "SessionMedium"
)

type AggregatedMetrics struct {
	Name                      string // Common field for Country, DeviceCategory, LandingPage, or SessionMedium
	AverageEngagementRate     float64
	TotalSessions             int
	TotalPageViews            int
	AverageSessionDuration    float64
	BounceRate                float64
	TotalNewUsers             int
	TotalUsers                int
	AverageEngagementDuration float64
	DataPointCount            int
}

func AggregateMetricsByBreakdown(data []model.Insight, breakdown Breakdown, threshold int) []AggregatedMetrics {
	metricsMap := make(map[string][]model.Insight)

	// Aggregate data based on the breakdown (Country, DeviceCategory, LandingPage, SessionMedium)
	for _, insight := range data {
		var key string
		switch breakdown {
		case COUNTRY:
			key = insight.Country
		case DEVICE:
			key = insight.DeviceCategory
		case PAGE:
			key = insight.LandingPage
		case MEDIUM:
			key = insight.SessionMedium
		}
		metricsMap[key] = append(metricsMap[key], insight)
	}

	// Calculate metrics for each breakdown category that meets the threshold
	var aggregatedMetrics []AggregatedMetrics
	for name, insights := range metricsMap {
		if len(insights) >= threshold && name != NOTSET {
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

			aggregatedMetrics = append(aggregatedMetrics, AggregatedMetrics{
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

	// Sort by average engagement rate in descending order
	sort.Slice(aggregatedMetrics, func(i, j int) bool {
		return aggregatedMetrics[i].AverageEngagementRate > aggregatedMetrics[j].AverageEngagementRate
	})

	return aggregatedMetrics
}

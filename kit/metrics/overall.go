package metrics

import "data-insights/kit/model"

func CalculateOverallMetrics(data []model.Insight) model.OverallMetrics {
	var totalEngagementRate, totalSessionDuration float64
	var totalSessions, totalPageViews, totalNewUsers, totalUsers, singlePageSessions int

	for _, insight := range data {
		engagementRate := parseStringToFloat(insight.EngagementRate)
		totalEngagementRate += engagementRate * float64(insight.Sessions)
		totalSessionDuration += float64(insight.UserEngagementDuration)
		if insight.Sessions > 0 && insight.ScreenPageViews == 1 {
			singlePageSessions += insight.Sessions
		}
		totalPageViews += insight.ScreenPageViews
		totalSessions += insight.Sessions
		totalNewUsers += insight.NewUsers
		totalUsers += insight.TotalUsers
	}

	overallEngagementRate := totalEngagementRate / float64(totalSessions)
	averageSessionDuration := totalSessionDuration / float64(totalSessions)
	bounceRate := (float64(singlePageSessions) / float64(totalSessions)) * 100
	pagesPerSession := float64(totalPageViews) / float64(totalSessions)
	newUserPercentage := (float64(totalNewUsers) / float64(totalUsers)) * 100
	sessionPerUser := float64(totalSessions) / float64(totalUsers)

	return model.OverallMetrics{
		OverallEngagementRate:  overallEngagementRate,
		AverageSessionDuration: averageSessionDuration,
		BounceRate:             bounceRate,
		PagesPerSession:        pagesPerSession,
		NewUserPercentage:      newUserPercentage,
		SessionPerUser:         sessionPerUser,
	}
}

package metrics

import "data-insights/kit/common"

// CalculateOverallMetrics calculates and returns overall metrics across the entire dataset, including overall engagement
// rate, average session duration, bounce rate, pages per session, new user percentage, and sessions per user.
func CalculateOverallMetrics(data []common.Insight) common.OverallMetrics {
	var totalEngagementRate, totalSessionDuration float64
	var totalSessions, totalPageViews, totalNewUsers, totalUsers, singlePageSessions int

	// Sum up the relevant metrics across all data points
	for _, insight := range data {
		engagementRate := parseStringToFloat(insight.EngagementRate)
		totalEngagementRate += engagementRate * float64(insight.Sessions)
		totalSessionDuration += float64(insight.UserEngagementDuration)
		if insight.Sessions > 0 && insight.ScreenPageViews == singlePageView {
			singlePageSessions += insight.Sessions
		}
		totalPageViews += insight.ScreenPageViews
		totalSessions += insight.Sessions
		totalNewUsers += insight.NewUsers
		totalUsers += insight.TotalUsers
	}

	// Calculate overall metrics, ensuring no division by zero
	var overallEngagementRate, averageSessionDuration, bounceRate, pagesPerSession, newUserPercentage, sessionPerUser float64

	if totalSessions > 0 {
		overallEngagementRate = totalEngagementRate / float64(totalSessions)
		averageSessionDuration = totalSessionDuration / float64(totalSessions)
		bounceRate = (float64(singlePageSessions) / float64(totalSessions)) * 100
		pagesPerSession = float64(totalPageViews) / float64(totalSessions)
	}

	if totalUsers > 0 {
		newUserPercentage = (float64(totalNewUsers) / float64(totalUsers)) * 100
		sessionPerUser = float64(totalSessions) / float64(totalUsers)
	}

	return common.OverallMetrics{
		OverallEngagementRate:  overallEngagementRate,
		AverageSessionDuration: averageSessionDuration,
		BounceRate:             bounceRate,
		PagesPerSession:        pagesPerSession,
		NewUserPercentage:      newUserPercentage,
		SessionPerUser:         sessionPerUser,
	}
}

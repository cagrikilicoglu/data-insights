package metrics

import "data-insights/kit/common"

const (
	thresholdDataPointNumber int = 100
	topBottomCount int = 5
)

// CalculateKeyMetrics calculates and returns key metrics for the dataset,
// including overall metrics and top/bottom breakdowns by country, device, page, and medium.
func CalculateKeyMetrics(data []common.Insight) common.UserMetrics {
	// Calculate overall metrics
	overallMetrics := CalculateOverallMetrics(data)

	// Aggregate and sort metrics by country
	metricsByCountry := AggregateMetricsByBreakdown(data, common.COUNTRY, thresholdDataPointNumber)
	metricsByCountry.SortByField(common.AVGENGAGEMENTRATE, common.DESC)

	// Aggregate and sort metrics by device category
	metricsByDevices := AggregateMetricsByBreakdown(data, common.DEVICE, thresholdDataPointNumber)
	metricsByDevices.SortByField(common.BOUNCERATE, common.DESC)

	// Aggregate and sort metrics by page
	metricsByPages := AggregateMetricsByBreakdown(data, common.PAGE, thresholdDataPointNumber)
	metricsByPages.SortByField(common.TOTALSESSIONS, common.DESC)

	// Aggregate and sort metrics by session medium
	metricsByMedium := AggregateMetricsByBreakdown(data, common.MEDIUM, thresholdDataPointNumber)
	metricsByMedium.SortByField(common.AVGSESSIONDURATION, common.DESC)

	// Return the aggregated user metrics with top and bottom elements for each breakdown
	return common.UserMetrics{
		OverallMetrics:                         overallMetrics,
		Top5CountriesWithHighestEngagementRate: GetTopElements(metricsByCountry, topBottomCount),
		Top5CountriesWithLowestEngagementRate:  GetBottomElements(metricsByCountry, topBottomCount),
		BounceRatesByDevices:                   metricsByDevices,
		Top5PagesWithHighestNoOfSessions:       GetTopElements(metricsByPages, topBottomCount),
		Top5PagesWithLowestNoOfSessions:        GetBottomElements(metricsByPages, topBottomCount),
		AverageSessionDurationsByDevices:       metricsByMedium,
	}
}

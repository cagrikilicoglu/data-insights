package metrics

import (
	"data-insights/kit/model"
)

const thresholdDataPointNumber int = 100

func GetImportantMetrics(data []model.Insight) model.UserMetrics {
	overallMetrics := CalculateOverallMetrics(data)

	metricsByCountry := AggregateMetricsByBreakdown(data, model.COUNTRY, thresholdDataPointNumber)
	metricsByCountry.SortByField(model.AVGENGAGEMENTRATE, model.DESC)

	metricsByDevices := AggregateMetricsByBreakdown(data, model.DEVICE, thresholdDataPointNumber)
	metricsByDevices.SortByField(model.BOUNCERATE, model.DESC)

	metricsByPages := AggregateMetricsByBreakdown(data, model.PAGE, thresholdDataPointNumber)
	metricsByPages.SortByField(model.TOTALSESSIONS, model.DESC)

	metricsByMedium := AggregateMetricsByBreakdown(data, model.MEDIUM, thresholdDataPointNumber)
	metricsByMedium.SortByField(model.AVGSESSIONDURATION, model.DESC)

	return model.UserMetrics{
		OverallMetrics:                         overallMetrics,
		Top5CountriesWithHighestEngagementRate: GetTopElements(metricsByCountry, 5),
		Top5CountriesWithLowestEngagementRate:  GetBottomElements(metricsByCountry, 5),
		BounceRatesByDevices:                   metricsByDevices,
		Top5PagesWithHighestNoOfSessions:       GetTopElements(metricsByPages, 5),
		Top5PagesWithLowestNoOfSessions:        GetBottomElements(metricsByPages, 5),
		AverageSessionDurationsByDevices:       metricsByMedium,
	}
}

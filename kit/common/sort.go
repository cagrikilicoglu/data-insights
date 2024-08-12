package common

import "sort"

type SortOrder string

const (
	DESC SortOrder = "DESC"
	ASC  SortOrder = "ASC"
)

// SortByField sorts a slice of AggregatedMetrics based on the specified field and order.
func (metrics AggregatedMetricsList) SortByField(field Metric, order SortOrder) {
	sort.Slice(metrics, func(i, j int) bool {
		// Determine the comparison based on the field and order
		var less bool
		switch field {
		case NAME:
			less = metrics[i].Name < metrics[j].Name
		case AVGENGAGEMENTRATE:
			less = metrics[i].AverageEngagementRate < metrics[j].AverageEngagementRate
		case TOTALSESSIONS:
			less = metrics[i].TotalSessions < metrics[j].TotalSessions
		case TOTALPAGEVIEWS:
			less = metrics[i].TotalPageViews < metrics[j].TotalPageViews
		case AVGSESSIONDURATION:
			less = metrics[i].AverageSessionDuration < metrics[j].AverageSessionDuration
		case BOUNCERATE:
			less = metrics[i].BounceRate < metrics[j].BounceRate
		case TOTALNEWUSERS:
			less = metrics[i].TotalNewUsers < metrics[j].TotalNewUsers
		case TOTALUSERS:
			less = metrics[i].TotalUsers < metrics[j].TotalUsers
		case AVGENGAGEMENTDURATION:
			less = metrics[i].AverageEngagementDuration < metrics[j].AverageEngagementDuration
		case DATAPOINTCOUNT:
			less = metrics[i].DataPointCount < metrics[j].DataPointCount
		default:
			return false
		}

		// Reverse the comparison if the order is DESC
		if order == DESC {
			return !less
		}
		return less
	})
}

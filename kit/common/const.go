package common

const NOTSET string = "(not set)"

type Breakdown string

const (
	COUNTRY Breakdown = "Country"
	DEVICE  Breakdown = "DeviceCategory"
	PAGE    Breakdown = "LandingPage"
	MEDIUM  Breakdown = "SessionMedium"
)

type Metric string

const (
	NAME                  Metric = "Name"
	AVGENGAGEMENTRATE     Metric = "AverageEngagementRate"
	TOTALSESSIONS         Metric = "TotalSessions"
	TOTALPAGEVIEWS        Metric = "TotalPageViews"
	AVGSESSIONDURATION    Metric = "AverageSessionDuration"
	BOUNCERATE            Metric = "BounceRate"
	TOTALNEWUSERS         Metric = "TotalNewUsers"
	TOTALUSERS            Metric = "TotalUsers"
	AVGENGAGEMENTDURATION Metric = "AverageEngagementDuration"
	DATAPOINTCOUNT        Metric = "DataPointCount"
)

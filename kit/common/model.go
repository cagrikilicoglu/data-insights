package common

type Insight struct {
	Country                string `json:"Country"`
	DeviceCategory         string `json:"DeviceCategory"`
	EngagementRate         string `json:"EngagementRate"`
	LandingPage            string `json:"LandingPage"`
	NewUsers               int    `json:"NewUsers"`
	ScreenPageViews        int    `json:"ScreenPageViews"`
	SessionMedium          string `json:"SessionMedium"`
	Sessions               int    `json:"Sessions"`
	TotalUsers             int    `json:"TotalUsers"`
	UserEngagementDuration int    `json:"UserEngagementDuration"`
	Date                   string `json:"date"`
}

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

type AggregatedMetricsList []AggregatedMetrics

type OverallMetrics struct {
	OverallEngagementRate  float64
	AverageSessionDuration float64
	BounceRate             float64
	PagesPerSession        float64
	NewUserPercentage      float64
	SessionPerUser         float64
}

type UserMetrics struct {
	OverallMetrics                         OverallMetrics
	Top5CountriesWithHighestEngagementRate AggregatedMetricsList
	Top5CountriesWithLowestEngagementRate  AggregatedMetricsList
	BounceRatesByDevices                   AggregatedMetricsList
	Top5PagesWithHighestNoOfSessions       AggregatedMetricsList
	Top5PagesWithLowestNoOfSessions        AggregatedMetricsList
	AverageSessionDurationsByDevices       AggregatedMetricsList
}

type OverallMetricsWithInsight struct {
	OverallEngagementRate  string `json:"overall_engagement_rate"`
	AverageSessionDuration string `json:"average_session_duration"`
	BounceRate             string `json:"bounce_rate"`
	PagesPerSession        string `json:"pages_per_session"`
	NewUserPercentage      string `json:"new_user_percentage"`
	SessionPerUser         string `json:"session_per_user"`
	AIInsight              string `json:"ai_insight"`
}

type AggregatedMetric struct {
	Name                   string `json:"name"`
	AverageEngagementRate  string `json:"average_engagement_rate,omitempty"`
	BounceRate             string `json:"bounce_rate,omitempty"`
	TotalSessions          string `json:"total_sessions,omitempty"`
	AverageSessionDuration string `json:"average_session_duration,omitempty"`
}

type AggregatedMetricsWithInsight struct {
	AIInsight         string             `json:"ai_insight"`
	AggregatedMetrics []AggregatedMetric `json:"aggregated_metrics"`
}

type UserMetricsWithInsights struct {
	OverallMetrics                         OverallMetricsWithInsight    `json:"overall_metrics"`
	Top5CountriesWithHighestEngagementRate AggregatedMetricsWithInsight `json:"top_5_countries_with_highest_engagement_rate"`
	Top5CountriesWithLowestEngagementRate  AggregatedMetricsWithInsight `json:"top_5_countries_with_lowest_engagement_rate"`
	BounceRatesByDevices                   AggregatedMetricsWithInsight `json:"bounce_rates_by_devices"`
	Top5PagesWithHighestNoOfSessions       AggregatedMetricsWithInsight `json:"top_5_pages_with_highest_no_of_sessions"`
	Top5PagesWithLowestNoOfSessions        AggregatedMetricsWithInsight `json:"top_5_pages_with_lowest_no_of_sessions"`
	AverageSessionDurationsByDevices       AggregatedMetricsWithInsight `json:"average_session_durations_by_devices"`
}

type EmailData struct {
	RecipientName string
	UserMetricsWithInsights
}

type EnvVariables struct {
	FileDirectory string
	ApiKey        string
	EmailFrom     string
	EmailPass     string
	EmailTo       string
	RecipientName string
	SmtpHost      string
	SmtpPort      string
}

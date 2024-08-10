package model

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

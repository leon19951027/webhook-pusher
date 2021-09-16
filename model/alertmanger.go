package model

type AlertManagerRequestBody struct {
	Alerts []struct {
		Annotations struct {
			Description string `json:"description"`
			Summary     string `json:"summary"`
		} `json:"annotations"`
		EndsAt       string `json:"endsAt"`
		Fingerprint  string `json:"fingerprint"`
		GeneratorURL string `json:"generatorURL"`
		Labels       struct {
			Alertname string `json:"alertname"`
			Instance  string `json:"instance"`
			Job       string `json:"job"`
			Severity  string `json:"severity"`
		} `json:"labels"`
		StartsAt string `json:"startsAt"`
		Status   string `json:"status"`
	} `json:"alerts"`
	CommonAnnotations struct {
		Description string `json:"description"`
		Summary     string `json:"summary"`
	} `json:"commonAnnotations"`
	CommonLabels struct {
		Alertname string `json:"alertname"`
		Instance  string `json:"instance"`
		Job       string `json:"job"`
		Severity  string `json:"severity"`
	} `json:"commonLabels"`
	ExternalURL string `json:"externalURL"`
	GroupKey    string `json:"groupKey"`
	GroupLabels struct {
		Alertname string `json:"alertname"`
	} `json:"groupLabels"`
	Receiver        string `json:"receiver"`
	Status          string `json:"status"`
	TruncatedAlerts int64  `json:"truncatedAlerts"`
	Version         string `json:"version"`
}

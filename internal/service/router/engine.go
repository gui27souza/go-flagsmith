package router

type DecideReq struct {
	UserID     string `json:"user_id"`
	Country    string `json:"country"`
	AppVersion string `json:"app_version"`
}

type TelemetryData struct {
	EvaluatedAt   time.Time `json:"evaluated_at"`
	CacheHydrated bool      `json:"cache_hydrated"`
}

type DecideRes struct {
	Target    string        `json:"target"`
	Reason    string        `json:"reason"`
	Telemetry TelemetryData `json:"telemetry"`
}


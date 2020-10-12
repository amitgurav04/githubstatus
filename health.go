package main

import "time"

type HealthStatus struct {
	Health    string    `json:"Github_health"`
	UpdatedAt time.Time `json:"Updated_at"`
}
type PageStruct struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	TimeZone  string    `json:"time_zone"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Status struct {
	Indicator   string `json:"indicator"`
	Description string `json:"description"`
}

type GitHubStatus struct {
	Page   PageStruct `json:"page"`
	Status Status     `json:"status"`
}

package main

import (
	"encoding/json"
)

func parseGithubStatus(body []byte) (*GitHubStatus, error) {
	var s = new(GitHubStatus)
	err := json.Unmarshal(body, &s)
	if err != nil {
		return nil, err
	}
	return  s, nil
}

func checkHealth(s *GitHubStatus) *HealthStatus {
	var health = new(HealthStatus)
	health.Health = s.Status.Indicator
	health.UpdatedAt = s.Page.UpdatedAt
	return health
}

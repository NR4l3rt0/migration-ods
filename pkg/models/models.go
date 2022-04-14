package models

import "time"

type ResultSet struct {
	ProjectKey          string     `json:"project_key"`
	RepositoryName      string     `json:"repository_name"`
	Technology          string     `json:"technology"`
	ComponentType       string     `json:"component_type"`
	ODSVersion          string     `json:"ods_version"`
	LastWebhookSuccess  *time.Time `json:"last_webhook_success"`
	ClusterRegion       string     `json:"cluster_region"`
	BuildName           string     `json:"build_name"`
	Environment         string     `json:"environment"`
	StartTimestamp      time.Time  `json:"start_timestamp"`
	CompletionTimestamp time.Time  `json:"completion_timestamp"`
	Duration            float64    `json:"duration"`
	Phase               string     `json:"phase"`
}

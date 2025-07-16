package model

import "time"

// Job 岗位模型
type Job struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Company     string    `json:"company"`
	Description string    `json:"description"`
	Requirement string    `json:"requirement"`
	JobCategory string    `json:"job_category"`
	CityInfo    string    `json:"city_info"`
	RecruitType string    `json:"recruit_type"`
	PublishTime time.Time `json:"publish_time"`
	Code        string    `json:"code"`
	CityList    []string  `json:"city_list,omitempty"`
}

package model

// Job 岗位模型
type Job struct {
    ID      int    `json:"id"`
    Title   string `json:"title"`
    Company string `json:"company"`
}
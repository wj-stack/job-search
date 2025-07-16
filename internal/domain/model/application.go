package model

// Application 投递记录模型
type Application struct {
    ID        int    `json:"id"`
    UserID    int    `json:"user_id"`
    JobID     int    `json:"job_id"`
    ResumeURL string `json:"resume_url"`
    Status    string `json:"status"`
}
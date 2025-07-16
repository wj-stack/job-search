package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/cron/v3"
)

// Job 定义岗位结构体
type Job struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Company string `json:"company"`
	Type    string `json:"type"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/job_search")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
}

// fetchJobsFromThirdParty 模拟从第三方平台获取岗位信息
func fetchJobsFromThirdParty() {
	// 模拟从第三方平台获取的岗位数据
	jobs := []Job{
		{Title: "后端开发工程师", Company: "XX 公司", Type: "全职"},
		{Title: "前端开发工程师", Company: "YY 公司", Type: "实习"},
	}

	for _, job := range jobs {
		_, err := db.ExecContext(context.Background(), "INSERT INTO jobs (title, company, job_type) VALUES (?, ?, ?)", job.Title, job.Company, job.Type)
		if err != nil {
			fmt.Printf("插入岗位信息失败: %v\n", err)
		}
	}
}

// searchJobs 实现岗位搜索接口
func searchJobs(c *gin.Context) {
	keyword := c.Query("keyword")
	jobType := c.Query("job_type")

	query := "SELECT id, title, company, job_type FROM jobs WHERE 1=1"
	args := []interface{}{}

	if keyword != "" {
		query += " AND title LIKE ?"
		args = append(args, "%"+keyword+"%")
	}
	if jobType != "" {
		query += " AND job_type = ?"
		args = append(args, jobType)
	}

	rows, err := db.QueryContext(c, query, args...)
	if err != nil {
		c.JSON(500, gin.H{"error": "数据库查询失败"})
		return
	}
	defer rows.Close()

	var result []Job
	for rows.Next() {
		var job Job
		if err := rows.Scan(&job.ID, &job.Title, &job.Company, &job.Type); err != nil {
			c.JSON(500, gin.H{"error": "数据解析失败"})
			return
		}
		result = append(result, job)
	}

	if err := rows.Err(); err != nil {
		c.JSON(500, gin.H{"error": "数据库查询出错"})
		return
	}

	c.JSON(200, result)
}

// 避免重复声明，直接将逻辑放在当前文件唯一的main函数中
// 初始化定时任务
// 初始化 Gin 路由
// 启动服务器

func main() {
	// 初始化定时任务
	c := cron.New()
	c.AddFunc("0 0 0 * * *", fetchJobsFromThirdParty) // 每天凌晨执行一次
	c.Start()

	// 初始化 Gin 路由
	r := gin.Default()
	r.GET("/api/jobs/search", searchJobs)

	// 启动服务器
	r.Run(":8080")
}

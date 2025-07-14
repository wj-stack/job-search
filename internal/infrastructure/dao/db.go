package dao

import (
	"entgo.io/ent/dialect/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wj-stack/job-search/ent"
)

// InitDB 初始化数据库连接
func InitDB(driver, dsn string) *ent.Client {
	drv, err := sql.Open(
		driver,
		dsn,
	)
	client := ent.NewClient(ent.Driver(drv))
	if err != nil {
		panic(err)
	}
	return client
}

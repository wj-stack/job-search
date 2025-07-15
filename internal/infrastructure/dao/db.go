package dao

import (
	"context"
	"log"

	"entgo.io/ent/dialect/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wj-stack/job-search/ent"
	"github.com/wj-stack/job-search/ent/migrate"
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
	err = client.Schema.Create(
		context.Background(),
		migrate.WithForeignKeys(false), // Disable foreign keys.
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

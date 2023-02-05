package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"log"
	"os"
	"train-tiktok/common/dbutil"
	"train-tiktok/service/video/internal/config"
	"train-tiktok/service/video/models"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {

	logx.MustSetup(c.Log)
	// Gorm
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = c.Mysql.DataSource
	}

	_db, err := dbutil.New(dsn, os.Getenv("DEBUG"))
	if err != nil {
		log.Panicf("failed to connect to mysql: %v", err)
	}

	//	自动生成表结构
	if err := _db.AutoMigrate(models.Video{}); err != nil {
		log.Panicf("failed to autoMigrate: %v", err)
	}

	return &ServiceContext{
		Config: c,
		Db:     _db,
	}
}

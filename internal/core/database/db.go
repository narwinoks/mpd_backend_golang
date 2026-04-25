package database

import (
	"backend-app/config"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewDatabase establishes a connection to the primary database
func NewDatabase(cfg *config.Config) (*gorm.DB, error) {
	return connect(cfg.Database)
}

// NewTestDatabase establishes a connection to the testing database
func NewTestDatabase(cfg *config.Config) (*gorm.DB, error) {
	return connect(cfg.DatabaseTest)
}

func connect(dbCfg config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Jakarta",
		dbCfg.Host,
		dbCfg.User,
		dbCfg.Password,
		dbCfg.Name,
		dbCfg.Port,
		dbCfg.SSLMode,
	)

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		logrus.Errorf("failed to connect database: %v", err)
		return nil, err
	}

	filterCallback := func(db *gorm.DB) {
		if !db.DryRun && db.Error == nil {
			FilterProfile(db)
		}
	}

	db.Callback().Query().Before("gorm:query").Register("filter_profile", filterCallback)
	db.Callback().Update().Before("gorm:update").Register("filter_profile", filterCallback)
	db.Callback().Delete().Before("gorm:delete").Register("filter_profile", filterCallback)
	db.Callback().Row().Before("gorm:row").Register("filter_profile", filterCallback)
	db.Callback().Raw().Before("gorm:raw").Register("filter_profile", filterCallback)

	orderCallback := func(db *gorm.DB) {
		if !db.DryRun && db.Error == nil {
			if _, skip := db.Get("gorm:without_default_order"); !skip {
				DefaultOrder(db)
			}
		}
	}
	db.Callback().Query().Before("gorm:query").After("filter_profile").Register("default_order", orderCallback)

	logrus.Infof("Database connection established: %s", dbCfg.Name)
	return db, nil
}

package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jniltinho/ftpdadmin/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Service interface {
	Health() map[string]string
	DB() *gorm.DB
}

type service struct {
	db *gorm.DB
}

func New() Service {

	lv := logger.Error
	if config.Server.Mode == "debug" {
		lv = logger.Info
	}

	cfg := &gorm.Config{Logger: logger.Default.LogMode(lv)}
	dsn := config.Database.DSN

	db, err := gorm.Open(mysql.Open(dsn), cfg)
	if err != nil {
		config.Fatal(err.Error(), "Cannot connect to database")
		//log.Fatal().Err(err).Msg("Cannot connect to database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB.SetConnMaxLifetime(0)
	sqlDB.SetMaxIdleConns(config.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.Database.MaxOpenConns)

	s := &service{db: db}
	return s
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	sqlDB, err := s.db.DB()
	if err != nil {
		log.Fatal(err)
	}

	err = sqlDB.PingContext(ctx)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}

func (s *service) DB() *gorm.DB {
	return s.db
}

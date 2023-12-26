package database

import (
	"sync"

	"github.com/jniltinho/ftpdadmin/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DBInstance is a singleton DB instance
type DBInstance struct {
	initializer func() any
	instance    any
	once        sync.Once
}

var dbInstance *DBInstance

// Instance gets the singleton instance
func (i *DBInstance) Instance() any {
	i.once.Do(func() { i.instance = i.initializer() })
	return i.instance
}

func dbInit() any {

	lv := logger.Error
	if config.Server.Mode == "debug" {
		lv = logger.Info
	}

	cfg := &gorm.Config{Logger: logger.Default.LogMode(lv)}

	db, err := gorm.Open(mysql.Open(config.Database.DSN), cfg)
	if err != nil {
		config.Fatal(err.Error(), "Cannot connect to database")
	}

	stdDB, _ := db.DB()
	stdDB.SetMaxIdleConns(config.Database.MaxIdleConns)
	stdDB.SetMaxOpenConns(config.Database.MaxOpenConns)

	return db
}

// DB returns the database instance
func DB() *gorm.DB {
	return dbInstance.Instance().(*gorm.DB)
}

func init() {
	dbInstance = &DBInstance{initializer: dbInit}

	// Create Default Tables if not exists
	initTables()
}

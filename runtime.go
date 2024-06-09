package taskmate

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-openapi/errors"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Runtime struct {
	Db  *gorm.DB
	Cfg *viper.Viper
}

func NewRuntime(cfg *viper.Viper) *Runtime {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.GetString("DB_HOST"),
		cfg.GetString("DB_USER"),
		cfg.GetString("DB_PASSWORD"),
		cfg.GetString("DB_NAME"),
		cfg.GetString("DB_PORT"),
	)

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	gormConfig := &gorm.Config{
		// enhance performance config
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		Logger:                 dbLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		log.Panicf("Error connect to DB : %f", err)
	}

	rt := &Runtime{
		Db:  db,
		Cfg: cfg,
	}

	return rt
}

func (r *Runtime) DB() *gorm.DB {
	return r.Db
}

func (r *Runtime) Config() *viper.Viper {
	return r.Cfg
}

func (r *Runtime) RunMigration() {
	r.Db.AutoMigrate()
}

func (r *Runtime) SetError(code int, msg string, args ...interface{}) error {
	return errors.New(int32(code), msg)
}

func (r *Runtime) GetError(err error) errors.Error {
	if v, ok := err.(errors.Error); ok {
		return v
	}

	return errors.New(http.StatusInternalServerError, err.Error())
}

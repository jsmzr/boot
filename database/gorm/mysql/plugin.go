package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jsmzr/boot"
	"github.com/jsmzr/boot/database/gorm/mysql/db"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormMysqlPlugin struct {
}

const configPrefix = "boot.gorm."

var defaultConfig map[string]interface{} = map[string]interface{}{
	"enabled":   true,
	"order":     5,
	"port":      3306,
	"host":      "127.0.0.1",
	"charset":   "utf8",
	"parseTime": true,
	"loc":       "Local",
	// if value <= 0 than not set
	"pool.maxOpenConns":    0,
	"pool.maxIdleConns":    0,
	"pool.connMaxLifetime": 0,
	"pool.connMaxIdleTime": 0,
}

func (g *GormMysqlPlugin) Load() error {
	database := viper.GetString(configPrefix + "database")
	if database == "" {
		return errors.New("not found config [database]")
	}
	username := viper.GetString(configPrefix + "username")
	if username == "" {
		return errors.New("not found config [username]")
	}
	password := viper.GetString(configPrefix + "password")
	if password == "" {
		logrus.Warn("database password unset")
	}
	host := viper.GetString(configPrefix + "host")
	port := viper.GetInt(configPrefix + "port")
	charset := viper.GetString(configPrefix + "charset")
	parseTime := viper.GetBool(configPrefix + "parseTime")
	loc := viper.GetString(configPrefix + "loc")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s", username, password, host, port, database, charset, parseTime, loc)
	instance, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}))
	if err != nil {
		return err
	}
	sqlDB, err := instance.DB()
	if err != nil {
		return err
	}
	setPool(sqlDB)

	db.DB = instance
	return nil
}

func setPool(sqlDB *sql.DB) {
	maxOpenConns := viper.GetInt(configPrefix + "pool.maxOpenConns")
	if maxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(maxOpenConns)
	}
	maxIdleConns := viper.GetInt(configPrefix + "pool.maxIdleConns")
	if maxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(maxIdleConns)
	}
	connMaxLifetime := viper.GetInt(configPrefix + "pool.connMaxLifetime")
	if connMaxLifetime > 0 {
		sqlDB.SetConnMaxLifetime(time.Second * time.Duration(connMaxLifetime))
	}
	connMaxIdleTime := viper.GetInt(configPrefix + "pool.connMaxIdleTime")
	if connMaxIdleTime > 0 {
		sqlDB.SetConnMaxIdleTime(time.Second * time.Duration(connMaxIdleTime))
	}
}

func (g *GormMysqlPlugin) Order() int {
	return viper.GetInt(configPrefix + "order")
}

func (g *GormMysqlPlugin) Enabled() bool {
	return viper.GetBool(configPrefix + "enabled")
}

func init() {
	boot.InitDefaultConfig(configPrefix, defaultConfig)
	boot.RegisterPlugin("gorm-mysql", &GormMysqlPlugin{})
}

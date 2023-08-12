package mysql

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jsmzr/boot"
	"github.com/jsmzr/boot/database/mysql/db"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type MysqlPlugin struct{}

const configPrefix = "boot.database."

var defaultConfig map[string]interface{} = map[string]interface{}{
	"enabled": true,
	"order":   5,
	"host":    "127.0.0.1",
	"port":    3306,
	"charset": "utf8",
}

func (m *MysqlPlugin) Load() error {
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
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?chatset=%s", username, password, host, port, database, charset)
	instance, err := sql.Open("mysql", url)
	if err != nil {
		return err
	}
	db.DB = instance
	return nil
}

func (m *MysqlPlugin) Order() int {
	return viper.GetInt(configPrefix + "order")
}

func (m *MysqlPlugin) Enabled() bool {
	return viper.GetBool(configPrefix + "enabled")
}

func init() {
	boot.InitDefaultConfig(configPrefix, defaultConfig)
	boot.RegisterPlugin("mysql", &MysqlPlugin{})
}

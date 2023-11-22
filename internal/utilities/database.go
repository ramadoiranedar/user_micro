package utilities

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetGormDbMysql(config *viper.Viper, log logrus.FieldLogger) (*gorm.DB, error) {
	CreateLogFields(logrus.StandardLogger(), TraceLog()).Info("setup gorm database mysql")

	dsn := GetGormDsnMysql(config)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

func GetGormDsnMysql(config *viper.Viper) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetString("db.user"),
		config.GetString("db.pass"),
		config.GetString("db.host"),
		config.GetString("db.port"),
		config.GetString("db.name"),
	)
}

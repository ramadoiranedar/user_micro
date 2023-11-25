package utilities

import (
	"fmt"

	"github.com/ramadoiranedar/user_micro/internal/constants"
	str "github.com/ramadoiranedar/user_micro/internal/constants/strings"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

func GetGormDbMysql(config *viper.Viper, logrus logrus.FieldLogger) (*gorm.DB, error) {
	EntryLogTrace(logrus, TraceLog()).Info("setup gorm database mysql")

	dsn := GetGormDsnMysql(config)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gorm_logger.Default.LogMode(gorm_logger.Info),
	})
}

func GetGormDsnMysql(config *viper.Viper) string {
	c := constants.NewConstants(config)
	return fmt.Sprintf(
		str.DSN_MYSQL_FORMAT,
		c.GetDbUser(),
		c.GetDbPass(),
		c.GetDbHost(),
		c.GetDbPort(),
		c.GetDbName(),
	)
}

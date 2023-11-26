package constants

import str "github.com/ramadoiranedar/user_micro/internal/constants/strings"

type constantsApp interface {
	GetAppName() string
	GetAppEnvironment() string
	GetAppMaxUploadMB() int64
}

func (c *constant) GetAppName() string {
	return c.config.GetString(str.CONFIG_APP_NAME)
}

func (c *constant) GetAppEnvironment() string {
	return c.config.GetString(str.CONFIG_APP_ENVIRONMENT)
}

func (c *constant) GetAppMaxUploadMB() int64 {
	return c.config.GetInt64(str.CONFIG_APP_MAX_UPLOAD_MB)
}

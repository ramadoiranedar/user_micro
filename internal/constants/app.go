package constants

import str "github.com/ramadoiranedar/user_micro/internal/constants/strings"

type constantsApp interface {
	GetAppVersion() string
	GetAppName() string
	GetAppDescription() string
	GetAppCopyright() string
	GetAppEnvironment() string
	GetAppMaxUploadMB() int64
}

func (c *constant) GetAppVersion() string {
	return c.config.GetString(str.CONFIG_APP_VERSION)
}

func (c *constant) GetAppName() string {
	return c.config.GetString(str.CONFIG_APP_NAME)
}

func (c *constant) GetAppDescription() string {
	return c.config.GetString(str.CONFIG_APP_DESCRIPTION)
}

func (c *constant) GetAppCopyright() string {
	return c.config.GetString(str.CONFIG_APP_COPYRIGHT)
}

func (c *constant) GetAppEnvironment() string {
	return c.config.GetString(str.CONFIG_APP_ENVIRONMENT)
}

func (c *constant) GetAppMaxUploadMB() int64 {
	return c.config.GetInt64(str.CONFIG_APP_MAX_UPLOAD_MB)
}

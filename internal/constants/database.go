package constants

import str "github.com/ramadoiranedar/user_micro/internal/constants/strings"

type constantsDb interface {
	GetDbUser() string
	GetDbPass() string
	GetDbHost() string
	GetDbPort() string
	GetDbName() string
}

func (c *constant) GetDbUser() string {
	return c.config.GetString(str.CONFIG_DB_USER)
}

func (c *constant) GetDbPass() string {
	return c.config.GetString(str.CONFIG_DB_PASS)
}

func (c *constant) GetDbHost() string {
	return c.config.GetString(str.CONFIG_DB_HOST)
}

func (c *constant) GetDbPort() string {
	return c.config.GetString(str.CONFIG_DB_PORT)
}

func (c *constant) GetDbName() string {
	return c.config.GetString(str.CONFIG_DB_NAME)
}

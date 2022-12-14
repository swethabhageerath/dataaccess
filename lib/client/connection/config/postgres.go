package config

import (
	"context"
	"fmt"

	"github.com/swethabhageerath/logging/lib/constants"
	"github.com/swethabhageerath/logging/lib/logger"
	"github.com/swethabhageerath/utilities/lib/utilities/helpers"
)

type PgConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func (p *PgConfig) Set(ctx context.Context, env helpers.IEnvironmentHelper, logger logger.ILogger) {
	p.host = env.Get("PG_HOST")
	p.port = env.Get("PG_PORT")
	p.user = env.Get("PG_USER")
	p.password = env.Get("PG_PW")
	p.dbname = env.Get("PG_DB")

	if p.host == "" || p.port == "" || p.user == "" || p.password == "" || p.dbname == "" {
		logger.AddContext(ctx).AddMessage("PostGres.Set() -> Database configuration is not provided").
			AddAppName("dataaccess").AddLogLevel(constants.ERROR).Log()
		panic("Database configuration is not provided")
	}
}

func (p *PgConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		p.host, p.port, p.user, p.password, p.dbname)
}

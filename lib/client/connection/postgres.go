package connection

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/swethabhageerath/dataaccess/lib/client/connection/config"
	"github.com/swethabhageerath/logging/lib/constants"
	"github.com/swethabhageerath/logging/lib/logger"
	"github.com/swethabhageerath/utilities/lib/utilities/helpers"
)

type Postgres struct {
	config config.PgConfig
	logger logger.ILogger
	env    helpers.IEnvironmentHelper
}

func NewPostgresConnection(config config.PgConfig, logger logger.ILogger, env helpers.IEnvironmentHelper) Postgres {
	return Postgres{
		config: config,
		logger: logger,
		env:    env,
	}
}

func (p Postgres) Connect(ctx context.Context) (*sql.DB, error) {
	p.config.Set(ctx, p.env, p.logger)
	db, err := sql.Open("postgres", p.config.String())
	if err != nil {
		p.logger.AddAppName("dataaccess").AddMessage(fmt.Sprintf("Postgres.Connect() -> Open() -> %s", err.Error())).
			AddLogLevel(constants.ERROR).AddFrames(constants.ALL).Log()
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		p.logger.AddAppName("dataaccess").AddMessage(fmt.Sprintf("Postgres.Connect() -> Ping() -> %s", err.Error())).
			AddLogLevel(constants.ERROR).AddFrames(constants.ALL).Log()

		return nil, err
	}

	return db, nil
}

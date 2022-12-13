package connection

import (
	"database/sql"
	"fmt"
	"libraries/dataaccess/lib/client/connection/config"
	"libraries/logging/lib/constants"
	"libraries/logging/lib/logger"
)

type Postgres struct {
	config config.PgConfig
	logger logger.ILogger
}

func NewPostgresConnection(config config.PgConfig, logger logger.ILogger) Postgres {
	return Postgres{
		config: config,
		logger: logger,
	}
}

func (p Postgres) Connect() (*sql.DB, error) {
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

package config

import (
	"github.com/timmbarton/layout/components/grpcserver"
	"github.com/timmbarton/layout/components/httpserver"
	"github.com/timmbarton/layout/components/postgresconn"
	"github.com/timmbarton/layout/components/redisconn"
	"github.com/timmbarton/layout/components/tracingconn"

	"github.com/timmbarton/example/internal/usecase"
)

type Config struct {
	HTTP     httpserver.Config
	GRPC     grpcserver.DefaultServerConfig
	Postgres postgresconn.Config
	Redis    redisconn.Config
	Tracing  tracingconn.Config
	UseCase  usecase.Config
}

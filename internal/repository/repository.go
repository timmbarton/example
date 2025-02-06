package repository

import (
	"github.com/timmbarton/layout/components/postgresconn"
)

type Repositories struct {
	Example ExampleRepository
}

func New(pg *postgresconn.Conn) *Repositories {
	return &Repositories{
		Example: &exampleRepository{
			pg: pg,
		},
	}
}

type ExampleRepository interface {
}

package repository

import "github.com/timmbarton/layout/components/postgresconn"

type exampleRepository struct {
	pg *postgresconn.Conn
}

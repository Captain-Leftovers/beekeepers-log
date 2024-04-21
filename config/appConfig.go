package appConfig

import "github.com/Captain-Leftovers/beekeepers-log/internal/database"

type AppConfig struct {
	DBQueries *database.Queries
}

func NewAppConfig(dbq *database.Queries) *AppConfig {
	return &AppConfig{
		DBQueries: dbq,
	}
}

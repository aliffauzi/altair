package core

import (
	"database/sql"
	"time"
)

type AppLoader interface {
	Compile(configPath string) (AppConfig, error)
}

type DatabaseLoader interface {
	Compile(configPath string) (map[string]DatabaseConfig, error)
}

type DatabaseConfig interface {
	Driver() string
	DBMigrationSource() string
	DBHost() string
	DBPort() (int, error)
	DBUsername() string
	DBPassword() string
	DBDatabase() string
	DBConnectionMaxLifetime() (time.Duration, error)
	DBMaxIddleConn() (int, error)
	DBMaxOpenConn() (int, error)
}

type DatabaseBearer interface {
	Database(dbName string) (*sql.DB, error)
}

type AppConfig interface {
	PluginExists(pluginName string) bool
	Plugins() []string
}

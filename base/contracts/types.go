package contracts

import (
	"database/sql"
	"time"

	"github.com/golang-migrate/migrate/v4/database"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

type ServiceProvider interface {
	Boot()
	Register()
}

type Config interface {
	Env(envName string, defaultValue ...any) any
	Add(name string, configuration any)
	Get(path string, defaultValue ...any) any
	GetString(path string, defaultValue ...any) string
	GetInt(path string, defaultValue ...any) int
	GetBool(path string, defaultValue ...any) bool
}

type Server interface {
	Grpc() ServerHandle
	Http() ServerHandle
}

type ServerHandle interface {
	Start()
	Run(string)
	Handler()
}

type Commands interface {
	Signature() string
	Flags() []cli.Flag
	Handle(*cli.Context) error
}

type Cache interface {
	Push(string, any) error
	Retrieve(string) []string
	Remove(string, int)
	Pop(string) []string
	Get(string, any) any
	Has(string) bool
	Set(string, any, time.Duration) error
	Pull(string, any) any
	Add(string, any, time.Duration) bool
	Remember(string, time.Duration, func() any) (any, error)
	RememberForever(string, func() any) (any, error)
	Forever(string, any) bool
	Forget(string) bool
	Flush() bool
}

type DatabaseDriver interface {
	GetConnection() string
	GetSqlDb() *sql.DB
	GetDriver() (database.Driver, error)
	GetGormDialect() gorm.Dialector
	GetDsn() string
}
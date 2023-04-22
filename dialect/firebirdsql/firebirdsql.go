package firebirdsql

import (
	"time"

	"github.com/doug-martin/goqu/v9"
)

func DialectOptions() *goqu.SQLDialectOptions {
	opts := goqu.DefaultDialectOptions()

	opts.TimeFormat = time.DateTime

	return opts
}

func init() {
	goqu.RegisterDialect("firebirdsql", DialectOptions())
}

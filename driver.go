package taos_native

import (
	"database/sql"

	"github.com/taosdata/driver-go/v3/af"
)

type TaosDriver struct {
	cfg    Config
	taos   *sql.DB
	afConn *af.Connector
}

func NewDriver(host string, port int, user string, password string, dbName string, prec string) *TaosDriver {
	return &TaosDriver{
		cfg: NewConfig(WithHost(host), WithPort(port), WithUserName(user),
			WithPassword(password), WithDatabase(dbName), WithPrecision(prec)),
	}
}

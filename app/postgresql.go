package app

import (
	"errors"
	"fmt"
	"github.com/avtara/carehub/utils"
	"github.com/gearintellix/u2"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

func (cfg *App) InitPostgres() (err error) {
	sqlConn := `
		host=__host__
		port=__port__
		user=__user__
		password=__password__
		dbname=__name__
		sslmode=__sslMode__
		application_name=__appKey__
	`

	sqlConn = u2.Binding(sqlConn, map[string]string{
		"host":     utils.GetEnv("database.host", "127.0.0.1"),
		"port":     utils.GetEnv("database.port", "5432"),
		"user":     utils.GetEnv("database.user", "admin"),
		"password": utils.GetEnv("database.password", "!ssecretdong~123"),
		"name":     utils.GetEnv("database.name", "carehub"),
		"sslMode":  utils.GetEnv("database.sslMode", "disable"),
		"appKey":   utils.GetEnv("database.appKey", "carehub"),
	})

	db, err := sqlx.Connect(utils.GetEnv("database.engine", "postgres"), sqlConn)
	if err != nil {
		return errors.Join(err, errors.New(fmt.Sprintf("Failed connect to database %s",
			utils.GetEnv("database.name", "carehub"))))
	}

	db.SetConnMaxLifetime(time.Minute *
		time.Duration(utils.ToInt(utils.GetEnv("database.connection_lifetime", "15"), 15)))
	db.SetMaxIdleConns(int(utils.ToInt(utils.GetEnv("database.maximum_idle", "5"), 5)))
	db.SetMaxOpenConns(int(utils.ToInt(utils.GetEnv("database.connection_max_open", "0"), 0)))

	cfg.DB = db

	return nil
}

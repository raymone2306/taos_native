package taos_native

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/taosdata/driver-go/v3/af"
	_ "github.com/taosdata/driver-go/v3/taosSql"
)

func (td *TaosDriver) Connect() error {
	taosDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", td.cfg.UserName, td.cfg.Password,
		td.cfg.Host, td.cfg.Port, td.cfg.Database)
	var err error
	td.taos, err = sql.Open("taosSql", taosDSN)
	if err != nil {
		return err
	}
	fmt.Println("Connected to " + taosDSN + " successfully.")
	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", td.cfg.Database)
	_, err = td.taos.Exec(query)
	if err != nil {
		log.Fatalln("Failed to create database power, ErrMessage: " + err.Error())
	}
	query = fmt.Sprintf("USE %s;", td.cfg.Database)
	_, err = td.taos.Exec(query)
	if err != nil {
		log.Fatalln("Failed to use database power, ErrMessage: " + err.Error())
	}
	// ANCHOR: pool
	// SetMaxOpenConns sets the maximum number of open connections to the database. 0 means unlimited.
	td.taos.SetMaxOpenConns(0)
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	td.taos.SetMaxIdleConns(2)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	td.taos.SetConnMaxLifetime(0)
	// SetConnMaxIdleTime sets the maximum amount of time a connection may be idle.
	td.taos.SetConnMaxIdleTime(0)
	// ANCHOR_END: pool
	return nil
}

func (td *TaosDriver) Close() {
	td.taos.Close()
}

func (td *TaosDriver) AFConnect() error {
	var err error
	td.afConn, err = af.Open(td.cfg.Host, td.cfg.UserName, td.cfg.Password, "", td.cfg.Port)
	if err != nil {
		return fmt.Errorf("Failed to connect to host:%s:%d, ErrMessage: %v\n", td.cfg.Host, td.cfg.Port, err)
	}
	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", td.cfg.Database)
	_, err = td.afConn.Exec(query)
	if err != nil {
		log.Fatalln("Failed to create database power, ErrMessage: " + err.Error())
	}
	query = fmt.Sprintf("USE %s;", td.cfg.Database)
	_, err = td.afConn.Exec(query)
	if err != nil {
		log.Fatalln("Failed to use database power, ErrMessage: " + err.Error())
	}
	return nil
}

func (td *TaosDriver) AFClose() {
	td.afConn.Close()
}

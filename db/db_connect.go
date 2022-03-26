package db

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

var Session *r.Session

// Rethink database name
var RDB string

func init() {
	// fmt.Println("db_connect")

	var err error
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	// rDB = os.Getenv("RDB_DB")
	RDB = cfg.Section("DATABASE").Key("RDB_DB").String()
	Session, err = r.Connect(r.ConnectOpts{
		Address:  cfg.Section("DATABASE").Key("RDB_IP").String() + ":" + cfg.Section("DATABASE").Key("RDB_PORT").String(),
		Database: RDB,
	})
	if err != nil {
		fmt.Printf("ERROR: There is issue connecting with the database.\nPlease make sure databse is accessible to Khoji by making sure settings in\nconfig.ini are setup properly and the database server is up and running.\n\n")
		fmt.Println("ERROR DETAILS:", err)
		os.Exit(1)
		return
	}
}

// Copyright © 2018-2020 Satinderjit Singh.
//
// See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at
// the top-level directory of this distribution for the individual copyright
// holder information and the developer policies on copyright and licensing.
//
// Unless otherwise agreed in a custom licensing agreement, no part of the
// kmdgo software, including this file may be copied, modified, propagated.
// or distributed except according to the terms contained in the LICENSE file
//
// Removal or modification of this copyright notice is prohibited.

package db

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gopkg.in/ini.v1"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

// Rethink database name
var rDB string

func init() {
	// fmt.Println("db")

	var err error
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	rDB = cfg.Section("DATABASE").Key("RDB_DB").String()
}

func CreateDb() {
	if rDB == "" {
		fmt.Println("Please select dbname")
		return
	}

	// dropDB(rDB)

	createDb, _ := r.DBCreate(rDB).Run(Session)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(createDb)
	if createDb == nil {
		fmt.Println("Database already exists:", rDB)
		return
	}
	if getObj(createDb) == "{}" {
		log.Println("Database created:", rDB)
	}

	// res, err := r.DB(rDB).Table("network").Changes().Run(Session)

	createTable(`blocks`, `hash`)
	createIndex(`blocks`, `height`)
	createIndex(`blocks`, `timestamp`)
	createIndex(`blocks`, `time`)
	createIndex(`blocks`, `difficulty`)
	createIndex(`blocks`, `miner`)
	createTable(`transactions`, `hash`)
	createIndex(`transactions`, `value`)
	createIndex(`transactions`, `timestamp`)
	createIndex(`transactions`, `blockHeight`)
	createIndex(`transactions`, `blockHash`)
	createIndex(`transactions`, `shieldedValue`)
	createTable(`accounts`, `address`)
	createIndex(`accounts`, `lastSeen`)
	createIndex(`accounts`, `firstSeen`)
	createIndex(`accounts`, `balance`)
	createTable(`network`, `name`)
	createTable(`logs`, ``)
	createTable(`stats`, `name`)
	createTable(`identities`, `name`)
	createIndex(`identities`, `identityaddress`)
	createIndex(`identities`, `parent`)
	createIndex(`identities`, `privateaddress`)
	createIndex(`identities`, `blockheight`)
	createTable(`sharedvout`, `hash`)
	createIndex(`sharedvout`, `hashvout`)
}

func DropDB() {
	_, err := r.DBDrop(rDB).Run(Session)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println("Database deleted:", rDB)
	// printStr("*** DB Drop result: ***")
	// printObj(result)
	// printStr("\n")
}

func createTable(table, _primaryKey string) {
	_, err := r.DB(rDB).TableCreate(table, r.TableCreateOpts{PrimaryKey: _primaryKey}).RunWrite(Session)
	if err != nil {
		fmt.Println(err)
	}

	log.Println("Table created:", table)
	// printStr("*** Create table result: ***")
	// printObj(result)
	// printStr("\n")
}

func createIndex(table, index string) {
	_, err := r.DB(rDB).Table(table).IndexCreate(index).RunWrite(Session)
	if err != nil {
		fmt.Println(err)
	}

	log.Printf("Index created for table - %v: %v\n", table, index)
	// printStr("*** Create table result: ***")
	// printObj(result)
	// printStr("\n")
}

// func printStr(v string) {
// 	fmt.Println(v)
// }

// func printObj(v interface{}) {
// 	vBytes, _ := json.Marshal(v)
// 	fmt.Println(string(vBytes))
// }

func getObj(v interface{}) string {
	vBytes, _ := json.Marshal(v)
	return string(vBytes)
}

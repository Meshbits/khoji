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

package createdb

import (
	"encoding/json"
	"flag"
	"fmt"

	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

var session *r.Session

// Rethink database name
var rDB string

func init() {
	var err error
	session, err = r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: rDB,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CreateDb(dnname string) {

	// rDBName := flag.String("dbname", "", "Rethink database name")
	// flag.Parse()
	// fmt.Println("dbname:", *rDBName)
	// rDB = *rDBName
	rDB = dnname

	if rDB == "" {
		fmt.Println("Please select dbname")
		flag.PrintDefaults()
		return
	}

	dropDB(rDB)

	createDb, err := r.DBCreate(rDB).Run(session)
	fmt.Println(err)
	fmt.Println(createDb)

	// res, err := r.DB(rDB).Table("network").Changes().Run(session)

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
	createTable(`sharedvout`, `hash`)
	createIndex(`sharedvout`, `hashvout`)
}

func dropDB(db string) {
	result, err := r.DBDrop(db).Run(session)
	if err != nil {
		fmt.Println(err)
		return
	}

	printStr("*** DB Drop result: ***")
	printObj(result)
	printStr("\n")
}

func createTable(table, _primaryKey string) {
	result, err := r.DB(rDB).TableCreate(table, r.TableCreateOpts{PrimaryKey: _primaryKey}).RunWrite(session)
	if err != nil {
		fmt.Println(err)
	}

	printStr("*** Create table result: ***")
	printObj(result)
	printStr("\n")
}

func createIndex(table, index string) {
	result, err := r.DB(rDB).Table(table).IndexCreate(index).RunWrite(session)
	if err != nil {
		fmt.Println(err)
	}

	printStr("*** Create table result: ***")
	printObj(result)
	printStr("\n")
}

func printStr(v string) {
	fmt.Println(v)
}

func printObj(v interface{}) {
	vBytes, _ := json.Marshal(v)
	fmt.Println(string(vBytes))
}

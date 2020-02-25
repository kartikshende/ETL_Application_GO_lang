package dbconn

import (
	"fmt"
	"log"
	"os"

	ent "../entities"
	pg "github.com/go-pg/pg"
)

func GetSQLDB() *pg.DB {

	//db connection using go-pg

	opts := &pg.Options{
		User:     "kartik",
		Password: "",
		Database: "etldb",
		//Addr:"http://127.0.0.1:51075",
	}

	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		log.Println("Failed to connect to database.\n")
		os.Exit(100)
	}

	fmt.Println("Connection Successfully")
	//ent.createFileinfoTable(db)
	ent.CreateFileinfoTable(db)
	ent.CreateTestsTable(db)

	return db
}

// dbDriver := "mysql"
// 	dbUser := "root"
// 	dbPass := ""
// 	dbName := "etldb"
// 	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()
// 	fmt.Println("Connection Successfully")

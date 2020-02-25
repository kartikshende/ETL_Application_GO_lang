package entities

import (
	"log"

	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

//Fileinfos struct to model the table
type Fileinfos struct {
	Filename string
}

//Savefile calls from main func saveFilename and insert value to databse
func (fi *Fileinfos) Savefile(db *pg.DB) error {

	insertErr := db.Insert(fi)
	if insertErr != nil {
		log.Println("Failed to insert fileinfo table:", insertErr)
		return insertErr
	}
	log.Println("File %s inserted successfully :", fi.Filename)
	return nil
}

//CreateFileinfoTable creates table if not exist
func CreateFileinfoTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&Fileinfos{}, opts)
	if createErr != nil {
		log.Println("Failed to create fileinfo table:", createErr)
		return createErr
	}
	log.Println("Fileinfo table created successfully")

	return nil

}

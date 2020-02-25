package entities

import (
	"log"

	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

type Tests struct {
	ID       int
	Name     string
	Lastname string
}

func CreateTestsTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&Tests{}, opts)
	if createErr != nil {
		log.Println("Failed to create tests table:", createErr)
		return createErr
	}
	log.Println("Tests table created successfully")

	return nil

}

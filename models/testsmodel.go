package models

import (
	ent "../entities"

	db "../dbconn"
)

type TestsModel struct {
}

// FindAll model tests and retrun [] to save
func (*TestsModel) FindAll() ([]ent.Tests, error) {
	db := db.GetSQLDB()
	var test []ent.Tests
	//_, err := db.Query(&test, "SELECT * FROM tests")
	err := db.Model(&test).Select()

	if err != nil {
		return nil, err
	} else {
		//fmt.Println(test)
		// for _, tes := range test {
		// 	fmt.Println(tes.ID)

		// }
		return test, nil
	}
}

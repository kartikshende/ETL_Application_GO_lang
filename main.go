package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	controll "./controllers"
	db "./dbconn"
	ent "./entities"
	"github.com/go-pg/pg"
	_ "github.com/lib/pq"
)

func main() {

	//db connection using go-pg
	db := db.GetSQLDB()

	defer db.Close()

	//listing directory files

	var fileinfo = getFileName(db)
	fmt.Println(fileinfo)

	directoryCheck := directoryCheck(db, fileinfo)
	//go directoryCheck(db, fileinfo)
	fmt.Println(directoryCheck)

	http.HandleFunc("/", controll.Index)

	http.ListenAndServe(":8080", nil)

}

//Directory check
func directoryCheck(db *pg.DB, listFilenameDb []ent.Fileinfos) bool {

	files, err := ioutil.ReadDir("testdirectory")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		//fmt.Println(file.Name())

		// fileCheck check file already in database or not return true if not
		fileCheck := stringInSlice(file.Name(), listFilenameDb)
		//fmt.Println(fileCheck)

		//checking file already in database
		if fileCheck != true {
			fmt.Println(file.Name())

			//rows := readFile("testdirectory/" + file.Name())

			//createdtable := createTable(db, rows, file.Name())

			//copyCSVfile copy all contents to database
			copycsv := copyCSVFile(db, file.Name())

			if copycsv == true {

				saveFilename(db, file.Name())

			}

			//if createdtable == true {
			//}
		}
		fmt.Println("No new file in directory")
	}

	return true
}

//Insert file name into db

func saveFilename(dbref *pg.DB, filename string) {

	newFi := &ent.Fileinfos{
		Filename: filename,
	}

	newFi.Savefile(dbref)

}

//copy csv file to database

func copyCSVFile(db *pg.DB, filedirectory string) bool {
	var filename = filedirectory
	fmt.Println(filename)

	file, err := os.Open("testdirectory/" + filedirectory)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	fileread := bufio.NewReader(file)

	//var extension = filepath.Ext(filename)
	//var name = filename[0 : len(filename)-len(extension)]

	//_, err = db.CopyFrom(fileread, `COPY `+name+` FROM STDIN WITH CSV HEADER`)
	_, err = db.CopyFrom(fileread, `COPY tests FROM STDIN WITH CSV HEADER`)
	if err != nil {
		log.Fatalf("failed to copy file: %s", err)
	}
	return true

}

//get file name from db
func getFileName(db *pg.DB) []ent.Fileinfos {

	var fileinfo []ent.Fileinfos
	//var query string = `SELECT filename FROM fileinfo`

	//_, err := db.Query(&fileinfo, query)
	err := db.Model(&fileinfo).Select()

	if err != nil {

		log.Fatalln("Can't get a file names from table:", err.Error())
	}
	fmt.Println(fileinfo)
	return fileinfo
}

//stringInSlice checks the filename is in array or not
func stringInSlice(filedirectory string, listFilenameDb []ent.Fileinfos) bool {
	for _, btempfile := range listFilenameDb {
		if btempfile.Filename == filedirectory {
			return true
		}
	}
	return false
}

//readfile read the csv file and return the first row array
// func readFile(name string) []string {
// 	f, err := os.Open(name)
// 	if err != nil {
// 		log.Fatalf("Cannot open '%s': %s\n", name, err.Error())
// 	}
// 	defer f.Close()
// 	r := csv.NewReader(f)
// 	r.Comma = ','
// 	r.Comment = '#'
// 	rows, err := r.ReadAll()
// 	if err != nil {
// 		log.Fatalln("Cannot read CSV data:", err.Error())
// 	}
// 	return rows[0]
// }

//createTable creates the new table for every csv file

// func createTable(db *pg.DB, rows []string, filename string) bool {

// 	//extension takes  "hello.csv or hello.blahblah" to trip the extention of file
// 	var extension = filepath.Ext(filename)
// 	var name = filename[0 : len(filename)-len(extension)]

// 	//create new table if not exists query
// 	_, err := db.Exec("CREATE TABLE IF NOT EXISTS " + name + "()")

// 	if err != nil {
// 		log.Fatalln("Can't create a new table:", err.Error())
// 		return false
// 	}

// 	//loop through file first row to add colunm to table
// 	for _, row := range rows {
// 		fmt.Println(row)
// 		_, err := db.Exec("ALTER TABLE " + name + " ADD COLUMN " + row + " TEXT")
// 		if err != nil {
// 			log.Fatalln("Can't update a table:", err.Error())
// 			return false
// 		}
// 	}

// 	return true
// }

//******* getting filenames from database *****

// var book []Fileinfo
// errrr := db.Model(&book).Select()
// if errrr != nil {
// 	panic(errrr)
// }

// fmt.Println(book)

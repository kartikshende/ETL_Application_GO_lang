// insert, err := db.Query("INSERT INTO fileinfo VALUES ( 2, 'TEST.csv' )")

// // if there is an error inserting, handle it
// if err != nil {
// 	panic(err.Error())
// }
// // be careful deferring Queries if you are using transactions
// defer insert.Close()

// file, err := os.Open("testdirectory/" + file.Name())

// 			if err != nil {
// 				log.Fatalf("failed opening file: %s", err)
// 			}

// 			scanner := bufio.NewScanner(file)
// 			scanner.Split(bufio.ScanLines)
// 			var txtlines []string

// 			for scanner.Scan() {
// 				txtlines = append(txtlines, scanner.Text())
// 			}

// 			file.Close()

// 			for _, eachline := range txtlines {
// 				fmt.Println(eachline)
// 			}

// txn, err := db.Begin()
// 			if err != nil {
// 				panic(err)
// 			}

// 			if _, err := txn.Prepare(pq.CopyIn("test", "name")); err != nil {
// 				panic(err)
// 			}

// 	//inside csv
// 	file, err := os.Open("testdirectory/" + file.Name())

// 	if err != nil {
// 		log.Fatalf("failed opening file: %s", err)
// 	}
// 	r := csv.NewReader(file)
// 	r.Comma = ','
// 	r.Comment = '#'

// 	records, err := r.ReadAll()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(records)

// 	//db.Create(&records[0])
// 	for _, i := range records {
// 		for _, j := range i {
// 			fmt.Print(j, "\t")
// 		}
// 		fmt.Println()
// 	}

// id_convert, _ := strconv.Atoi(records[0][0])
// 			filetest := Test{
// 				ID:       id_convert,
// 				Name:     records[1][1],
// 				Lastname: records[2][2],
// 			}
// db.Create(&filetest)

// stmt, err := db.Prepare("INSERT INTO test(id, name, lastname) VALUES(?, ?, ?)")
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			for _, row := range rows[1:] {
// 				_, err := stmt.Exec(row[0], row[1], row[2])
// 				if err != nil {
// 					log.Fatal(err)
// 				}
// 			}

//********* Select query ******

// var book []Tests
// 	errrr := db.Model(&book).Select()
// 	if errrr != nil {
// 		panic(errrr)
// 	}
// 	fmt.Println(book)
// 	var filetests []string
// 	var query2 string = `SELECT * FROM tests`

// 	_, errr := db.Query(&filetests, query2)

// 	if errr != nil {
// 		fmt.Println("problem here")
// 		panic(errr.Error())
// 	}
// 	fmt.Println(filetests)

/*----------gorm-----------*/
//GetMySQLDB()
//Db connection using gorm
//db, err := gorm.Open("postgres", "user=kartik dbname=etldb sslmode=disable")

//db.SingularTable(true)
//db.CreateTable(&Test{})
//db.AutoMigrate(&Test{})
//results, err := db.Table("fileinfo").Select("filename").Rows()

/*----------sql lib/pg-----------*/

//connStr := "user=kartik dbname=etldb sslmode=disable"
//db, err := sql.Open("postgres", connStr)
// if err != nil {
// 	log.Fatal(err)
// }

//results, err := db.Query("SELECT filename FROM fileinfo")

//fmt.Println(result)

// names := make([]string, 0)

// for results.Next() {
// 	var name string
// 	if err := results.Scan(&name); err != nil {
// 		// Check for a scan error.
// 		// Query rows will be closed with defer.
// 		log.Fatal(err)
// 	}
// 	//fmt.Println(name)
// 	names = append(names, name)

// }
// fmt.Println(names)

//db connection using go-pg
// opts := &pg.Options{
// 	User:     "kartik",
// 	Password: "",
// 	Database: "etldb",
// 	//Addr:"http://127.0.0.1:51075",
// }

// db := pg.Connect(opts)

// defer db.Close()
// fmt.Println("Connection Successfully")

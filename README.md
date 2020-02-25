# ‘ETL’ (Extract, Transform, Load) tool in Go

### Pre-requiremt
* [go lang](https://golang.org/doc/) installed
* [postgres SQL](https://www.postgresql.org/docs/) installed

### Features of app

* This application perform ETL task whenver new(csv) file added in directory.
* The file(csv) is read and its content stored in a database if any new file detected.
* Database also stores meta information(filename) of any file, so what files already got processed in order to not process them again.


## Postgres SQL databse used to perform task
* Also used [go-pg](https://godoc.org/github.com/go-pg/pg#pkg-examples) orm which is very flexible and provide good fuctionalty to process data

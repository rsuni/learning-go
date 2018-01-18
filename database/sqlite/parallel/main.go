package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var wg sync.WaitGroup

func main() {

	os.Remove("./foo.db")

	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)
	defer db.Close()

	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	timeStart := time.Now()
	//tx, err := db.Begin()//without transaction database is even locked,
	//checkErr(err)

	//stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	stmt, err := db.Prepare("insert into foo(id, name) values(?, ?)")
	checkErr(err)
	defer stmt.Close()

	count := 100 //1000000 parallel inserts block whole computer
	wg.Add(count)
	for i := 0; i < count; i++ {
		//_, err = stmt.Exec(i, fmt.Sprintf("???????%03d", i))
		//checkErr(err)
		insert(stmt, i)
	}
	wg.Wait()
	//tx.Commit()
	timeEnd := time.Now()

	delta := timeEnd.Sub(timeStart)
	fmt.Println(delta.Seconds())

	fmt.Printf("%s - %s\n", timeStart, timeEnd)

	rows, err := db.Query("SELECT COUNT(*) as count FROM  foo")
	fmt.Println("Total count:", checkCount(rows))
	checkErr(err)

	return
}

func insert(stmt *sql.Stmt, i int) {
	defer wg.Done()
	_, err := stmt.Exec(i, fmt.Sprintf("???????%03d", i))

	checkErr(err)
}

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		checkErr(err)
	}
	return count
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

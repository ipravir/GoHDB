package main

import (
	"database/sql"
	"fmt"
	"log"

	// Register hdb driver.
	_ "github.com/SAP/go-hdb/driver"
)

const driverName = "hdb"

func getDbUrl() string {
	return "hdb" + "://" +
		"SYSTEM" + ":" + "Abap@2113" + "@" +
		"192.168.100.198" + ":" + "39013"
}

func main() {

	db, err := sql.Open(driverName, getDbUrl())
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Successfull Connected")

	// Fetching total number of records from emp_master table
	TestSchema := "SYSTEM"
	table := "emp_master"
	var i = 0
	if err := db.QueryRow(fmt.Sprintf("select count(*) from %s.%s", TestSchema, table)).Scan(&i); err != nil {
		log.Fatal(err)
	}
	fmt.Println(i)

	var result string
	if err := db.QueryRow(fmt.Sprintf("select fname from %s.%s where id = :1", TestSchema, table), 1).Scan(&result); err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	//rows, err := db.Query("SELECT id, fname FROM emp_master LIMIT $1", i)
	rows, err := db.Query("SELECT id, fname, lname FROM emp_master")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var fname string
		var lname string
		err = rows.Scan(&id, &fname, &lname)
		fmt.Println(id, fname, lname)
	}
}

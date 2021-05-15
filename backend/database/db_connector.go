package database

import (
	"database/sql"
	"fmt"
	"log"
)

func Connect() {
	// Create the database handle, confirm driver is present
	db, err := sql.Open("mysql", "fm:marmar@/test")
	if err != nil {
		log.Fatal(err)
	}
	defer fmt.Println("closing connection")
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("")
	}

	var (
		personID  int
		lastName  string
		firstName string
		address   string
		city      string
	)
	rows, err := db.Query("select * from Persons")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&personID, &lastName, &firstName, &address, &city)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(personID, lastName, firstName, firstName, address, city)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

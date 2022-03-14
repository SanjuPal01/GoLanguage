package util

import (
	"database/sql"
	"fmt"
)

func Insert(db *sql.DB) {
	sqlStatement := `INSERT INTO stu_details (id, firstname, lastname, age) 
    VALUES ($1, $2, $3, $4)`
	var id, fname, lname string
	var age int
	fmt.Println("Enter ID, FirstName, LastName and Age")
	fmt.Scanln(&id, &fname, &lname, &age)
	_, err := db.Exec(sqlStatement, id, fname, lname, age)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nRow inserted successfully!")
	}
}

package util

import (
	"database/sql"
	"fmt"
	"strconv"
)

func Insert(db *sql.DB) {
	sqlStatement := `INSERT INTO stu_details (id, firstname, lastname, age) 
    VALUES ($1, $2, $3, $4)`
	var id, fname, lname, age string
	fmt.Println()
	fmt.Println("----------------------------------------------------------------------------------------------------------")
	fmt.Print("Enter ID, FirstName, LastName and Age(In Respective Order): ")
	fmt.Scanln(&id, &fname, &lname, &age)
	ageNum, err := strconv.Atoi(age)
	if !validate(id, fname, lname, ageNum) {
		return
	} else if err != nil {
		fmt.Println("\nEnter Age in Valid Format")
		fmt.Println("----------------------------------------------------------------------------------------------------------")
		fmt.Println()
		return
	}

	_, err = db.Exec(sqlStatement, id, fname, lname, ageNum)
	if err != nil {
		fmt.Println("\n Student ID already Registered")
	} else {
		fmt.Println("\nRow inserted successfully!")
	}
	fmt.Println("----------------------------------------------------------------------------------------------------------")
	fmt.Println()
}

func validate(id, fname, lname string, ageNum int) bool {
	if len(id) == 0 || len(fname) == 0 || len(lname) == 0 {
		fmt.Println("\nPlease Enter All Details")
		fmt.Println("----------------------------------------------------------------------------------------------------------")
		fmt.Println()
		return false
	}
	if ageNum <= 0 {
		fmt.Println("\nEnter Age in Valid Format")
		fmt.Println("----------------------------------------------------------------------------------------------------------")
		fmt.Println()
		return false
	}
	return true
}

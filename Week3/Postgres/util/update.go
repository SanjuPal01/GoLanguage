package util

import (
	"database/sql"
	"fmt"
	"strconv"
)

func Update(db *sql.DB) {
	sqlStatement := `
		UPDATE stu_details
		SET firstname = $2, lastname = $3, age = $4
		WHERE id = $1;`
	var id, fname, lname, age string
	fmt.Println()
	fmt.Println("----------------------------------------------------------------------------------------------------------")
	fmt.Print("Please Enter Student ID: ")
	fmt.Scanln(&id)
	fmt.Print("Please Enter Updated Data(FirstName, LastName and Age): ")
	fmt.Scanln(&fname, &lname, &age)
	ageNum, err := strconv.Atoi(age)
	if !validate(id, fname, lname, ageNum) {
		return
	} else if err != nil {
		fmt.Println("\nEnter Age in Valid Format")
		fmt.Println("----------------------------------------------------------------------------------------------------------")
		fmt.Println()
		return
	}
	res, err := db.Exec(sqlStatement, id, fname, lname, ageNum)
	if err != nil {
		fmt.Println(err)
	}
	rowsAffected, _ := res.RowsAffected()
	if err != nil {
		fmt.Println("\n", err.Error())
	} else if rowsAffected == 0 {
		fmt.Println("\nEnter Valid ID")
	} else {
		fmt.Println("\nRecords Updated")
	}
	fmt.Println("----------------------------------------------------------------------------------------------------------")
	fmt.Println()
}

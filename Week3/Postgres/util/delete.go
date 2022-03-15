package util

import (
	"database/sql"
	"fmt"
)

func Delete(db *sql.DB) {
	sqlStatement := `
		DELETE FROM stu_details
		WHERE id = $1;`
	var id string
	fmt.Println()
	fmt.Println("------------------")
	fmt.Print("Enter Student ID: ")
	fmt.Scanln(&id)
	res, err := db.Exec(sqlStatement, id)
	rowsAffected, _ := res.RowsAffected()
	if err != nil {
		fmt.Println("\n", err.Error())
	} else if rowsAffected == 0 {
		fmt.Println("\nEnter Valid ID")
	} else {
		fmt.Println("\nRecords Deleted")
	}
	fmt.Println("------------------")
	fmt.Println()

}

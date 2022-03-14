package util

import (
	"database/sql"
	"fmt"
)

type Detail struct {
	id        string
	firstname string
	lastname  string
	age       int
}

func RetrieveAll(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM stu_details")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	studentDetail := make([]Detail, 0)
	for rows.Next() {
		tempStudent := Detail{}
		err := rows.Scan(&tempStudent.id, &tempStudent.firstname, &tempStudent.lastname, &tempStudent.age)
		if err != nil {
			panic(err)
		}
		studentDetail = append(studentDetail, tempStudent)
	}
	for _, val := range studentDetail {
		fmt.Println(val.id, val.firstname, val.lastname, val.age)
	}
}

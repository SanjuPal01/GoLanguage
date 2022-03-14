package skelton

import (
	"database/sql"
	"fmt"

	"example.com/util"
)

func Structure(db *sql.DB) {
	fmt.Println("---------------------")
	fmt.Println("Student Details Forum")
	fmt.Println("---------------------")
	for {
		fmt.Println("Enter 1 to Insert student Detail")
		fmt.Println("Enter 2 to Delete student Detail")
		fmt.Println("Enter 3 to Update student Detail")
		fmt.Println("Enter 4 to Show one student Detail")
		fmt.Println("Enter 5 to Show all student Detail")
		fmt.Println("Enter -1 to Exit")
		var val int
		fmt.Print("Enter Your Choice: ")
		fmt.Scanln(&val)
		switch val {
		case 1:
			fmt.Println("Insert")
			util.Insert(db)
		case 2:
			fmt.Println("Delete")
		case 3:
			fmt.Println("Update")
		case 4:
			fmt.Println("Show One")
		case 5:
			fmt.Println("Show All")
			util.RetrieveAll(db)
		case -1:
			fmt.Println("Show Exit")
			return
		default:
			fmt.Println("Enter Valid Input")
		}
	}
}

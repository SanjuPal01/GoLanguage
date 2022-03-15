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
		fmt.Println("Enter 4 to Show all student Detail")
		fmt.Println("Enter -1 to Exit")
		var val int
		fmt.Print("Enter Your Choice: ")
		fmt.Scanln(&val)
		switch val {
		case 1:
			util.Insert(db)
		case 2:
			util.Delete(db)
		case 3:
			util.Update(db)
		case 4:
			util.RetrieveAll(db)
		case -1:
			return
		default:
			fmt.Println("Please Enter Valid Input")
		}
	}
}

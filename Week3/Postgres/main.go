package main

import (
	"example.com/connection"
	"example.com/skelton"
)

func main() {
	db := connection.ConnectDB()
	skelton.Structure(db)
}

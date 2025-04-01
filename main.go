package main

import (
	"fmt"
	"ymoutella/ead/pkg/db"
)

func main() {

	var database db.Db
	fmt.Print(database.Connect())

}

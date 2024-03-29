package Controllers

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var dbShield *sql.DB

func CreateConnection() {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/tiffino")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	rows, _ := db.Query("SELECT id, name FROM users")
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}
}

func CreateShieldDb() {
	var err error
	dbShield, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/shield")
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = dbShield.Ping()
	if err != nil {
		log.Fatalf(err.Error())
	}

}

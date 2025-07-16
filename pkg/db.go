package pkg

import (
	_ "net/http"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gorilla/mux"
)

func NewConnection() *sql.DB{
	db, err := sql.Open("mysql", "root:pass@tcp(127.0.0.1:3306)/fake_tw_database")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	 err = db.Ping()
	 if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}

	return db
}
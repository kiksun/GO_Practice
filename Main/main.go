package main
import (
  "database/sql"
  _"github.com/go-sql-driver/mysql"
)
func main() {
  db, err := sql.Open("mysql", "root:FdVSTh9L@tcp(127.0.0.1:3306)/mydb")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()
  insert, err := db.Query("INSERT INTO posts(title) VALUES('AAA')")
  if err != nil {
    panic(err.Error())
  }
  defer insert.Close()
}
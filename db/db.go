package db

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

// InitDB - データベース接続の初期化
func InitDB() *sql.DB {
    db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
    if err != nil {
        panic(err)
    }
    return db
}

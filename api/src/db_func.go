package src

import (
   "database/sql"
   "fmt"
)

const (
    host = "db"
    port = "5432"
    user = "psql"
    password = "psql"
    dbname = "db"
)

func checkerr(err error){
    if err != nil{
        panic(err)
    }
}

func Opendb() (*sql.DB){
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)
    checkerr(err)
    err = db.Ping()
    checkerr(err)

    return db
}

func Countrows(db *sql.DB) (int){ 
    var res int
    err := db.QueryRow(`select count (*) from urls`).Scan(&res)
    checkerr(err)

    return res
}

func Insertdb(db *sql.DB, orig_url string, short_id string){
    sqlStatement := `insert into urls(orig_url, short_id) values ($1, $2)`
    _, err := db.Exec(sqlStatement, orig_url, short_id)
    checkerr(err)
}

func CheckExist(db *sql.DB, short_id string) (bool){
    var res int
    err := db.QueryRow(`select count (*) from urls where short_id=$1`, short_id).Scan(&res)
    checkerr(err)

    return (res > 0)
}

func Get_orig_url(db *sql.DB, short_id string) (string){
    var res string
    err := db.QueryRow(`select orig_url from urls where short_id=$1`, short_id).Scan(&res)
    checkerr(err)

    return res
}

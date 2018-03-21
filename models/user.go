package models

import (
	"github.com/astaxie/beego"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var err error

func init() {
	db, err = sql.Open("mysql", "demo:demo@/demo?charset=utf8")
	if err = db.Ping(); err != nil {
		beego.Info(err)
	} 
}

func FindUser(name string) string {
	var fullname string
	row := db.QueryRow("SELECT fullname FROM users WHERE username = '" + name + "'")
	err = row.Scan(&fullname)
	if err != nil {
		fullname = ""
	}
	return fullname
}

func VerifyUser(name string, password string) string {
	var fullname string
	var hash string
	row := db.QueryRow("SELECT fullname,password FROM users WHERE username = '" + name + "'")
	err = row.Scan(&fullname, &hash)
	if err != nil {
		return ""
	}
	// beego.Info("Hash: " + hash)
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return ""
	}	
	return fullname
}

func AddUser(username string, fullname string, password string) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	_ , err := db.Exec("INSERT INTO users (username, fullname, password) VALUES ('" + username + "', '" + fullname + "', '" + string(bytes) + "')")
	if err != nil {
		beego.Info(err)
	}
}
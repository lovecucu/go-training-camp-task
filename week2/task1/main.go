package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initMySQL() (err error) {
	dsn := "root:@tcp(10.24.16.12:3306)/user_0"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	return
}

type User struct {
	uid   int
	uname string
	level int
}

func (u User) String() string {
	return fmt.Sprintf("user.id is %d, user.uname is %s, user.level is %d\n", u.uid, u.uname, u.level)
}

func getUserNameByUid(uid int) (User, error) {
	user := User{uid: uid}
	err := db.QueryRow("SELECT uname,level FROM user_info_00 WHERE uid = ?", uid).Scan(&user.uname, &user.level)
	if errors.Is(err, sql.ErrNoRows) {
		err = nil // 直接忽略，也不记日志
	}
	return user, err
}

func main() {
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	uid := 1000000000
	user, err := getUserNameByUid(uid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(user)
}

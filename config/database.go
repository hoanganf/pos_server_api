package config

import (
	"database/sql"
	"gopkg.in/gorp.v1"
_ "github.com/go-sql-driver/mysql"
	"fmt"
	"os"
	"log"
)

func InitDB() *gorp.DbMap {
	user := getEnvWithDefault("DB_USER", "root")
	password := getEnvWithDefault("DB_PASSWORD", "")
//	host := getEnvWithDefault("DB_HOST", "127.0.0.1")
//	port := getEnvWithDefault("DB_PORT", "3306")
	dbName := getEnvWithDefault("DB_NAME", "anit_pos_server_new")
//	dsn := fmt.Sprintf("%s:%s@unix(%s:%s)/%s?parseTime=true", user, password, host, port,dbName)
	dsn := fmt.Sprintf("%s:%s@unix(/Applications/XAMPP/xamppfiles/var/mysql/mysql.sock)/%s?parseTime=true", user, password,dbName)
	fmt.Printf("dns: %s",dsn)
	db, err := sql.Open("mysql", dsn)
  checkErr(err, "sql.Open failed")

  return &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
}

func checkErr(err error, msg string) {
  if err != nil {
    log.Fatalln(msg, err)
  }
}

func getEnvWithDefault(name, def string) string {
	env := os.Getenv(name)
	if len(env) != 0 {
		return env
	}
	return def
}

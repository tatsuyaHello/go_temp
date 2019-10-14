package db

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var DBConn *gorm.DB

func init() {
	var err error
	dataSource := GetDataSource()
	DBConn, err = newMySQLConn(dataSource)
	if err != nil {
		log.Fatalf("mysqlのコネクション確立に失敗:%v", err)
	}
}

func newMySQLConn(dataSource string) (*gorm.DB, error) {
	return gorm.Open("mysql", dataSource)
}

type mySQLConfig struct {
	user     string
	password string
	host     string
	port     string
	dbName   string
}

func (c *mySQLConfig) dataSource() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.user,
		c.password,
		c.host,
		c.port,
		c.dbName,
	)
}

func GetDataSource() string {
	var err error
	dotenvPath := os.Getenv("APP_DOTENV_PATH")
	if dotenvPath == "" {
		log.Println("cannot find APP_DOTENV_PATH")
		err = godotenv.Load()
	} else {
		err = godotenv.Load(dotenvPath)
	}
	if err != nil {
		log.Fatalf(".envファイルの読み込みに失敗しました:%e", err)
	}

	config := mySQLConfig{
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		dbName:   os.Getenv("DB_NAME"),
	}
	dataSource := config.dataSource()
	return dataSource
}

package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

//ConnectDB Connexion à la BDD
func ConnectDB() {
	dsn := getDSN()
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Erreur survenue pendant la connexion à la DB: " + err.Error())
	}
	fmt.Println("Connexion à la BD réussie!")
	MigrateDatabase()
}

//getDSN DNS du sql
func getDSN() string {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	address := os.Getenv("DB_ADDRESS")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dns := user + ":" + password + "@tcp(" + address + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	return dns
}

package main

import (
	"Golang_Fiber/application"
	"Golang_Fiber/database"
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"strconv"
)

func main() {
	//Load env file
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Println("Erreur pendant le chargement du fichier env")
		panic(err.Error())
	}
	fresh, _ := strconv.ParseBool(os.Getenv("FRESH"))
	seed, _ := strconv.ParseBool(os.Getenv("SEED"))

	if fresh || seed {
		application.CliCommandApp(fresh, seed)
		os.Exit(0)
	}

	//Définir le fichier de sortie pour les logs
	log.SetOutput(&lumberjack.Logger{
		Filename:   "logs/server.log",
		MaxSize:    20, // megabytes
		MaxBackups: 3,
		MaxAge:     30, //days
	})

	defer database.CloseConnections()

	app := application.InitFiberApp()

	log.Panic(app.Listen(":3000"))
}

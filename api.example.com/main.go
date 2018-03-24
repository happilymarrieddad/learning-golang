package main

import (
	"./src/system/app"
	DB "learning-golang/api.example.com/src/system/db"

	"flag"
	"github.com/joho/godotenv"
	"os"
)

var port string
var dbhost string
var dbport string
var dbuser string
var dbpass string
var dboptions string
var dbdatabase string

func init() {
	flag.StringVar(&port, "port", "8000", "Assigning the port that the server should listen on.")
	flag.StringVar(&dbhost, "dbhost", "127.0.0.1", "Set the port for the application")
	flag.StringVar(&dbport, "dbport", "3306", "Set the port for the application")
	flag.StringVar(&dbuser, "dbuser", "root", "Set the port for the application")
	flag.StringVar(&dbpass, "dbpass", "pass", "Set the port for the application")
	flag.StringVar(&dboptions, "dboptions", "parseTime=true", "Set the port for the application")
	flag.StringVar(&dbdatabase, "dbdatabase", "fusion", "Set the port for the application")

	flag.Parse()

	if err := godotenv.Load("config.ini"); err != nil {
		panic(err)
	}

	if host := os.Getenv("DB_HOST"); len(host) > 0 {
		dbhost = host
	}
	if database := os.Getenv("DB_DATABASE"); len(database) > 0 {
		dbdatabase = database
	}
	if user := os.Getenv("DB_USER"); len(user) > 0 {
		dbuser = user
	}
	if password := os.Getenv("DB_PASSWORD"); len(password) > 0 {
		dbpass = password
	}
	if port := os.Getenv("DB_PORT"); len(port) > 0 {
		dbport = port
	}

	envPort := os.Getenv("PORT")
	if len(envPort) > 0 {
		port = envPort
	}
}

func main() {
	db, err := DB.Connect(dbhost, dbport, dbuser, dbpass, dbdatabase, dboptions)
	if err != nil {
		panic(err)
	}

	s := app.NewServer()

	s.Init(port, db)
	s.Start()
}

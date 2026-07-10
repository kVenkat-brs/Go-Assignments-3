package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kVenkat-brs/Go-Assignments-3/src"
)

func main() {

	if err:= godotenv.Load();err!=nil{
		log.Fatal("error loading Env's",err)
	}

	app := src.SetupApp()

	port := os.Getenv("PORT")

	if port == ""{
		port = ":3000"
	}


	fmt.Println("Server started on port",port)

	if err :=app.Listen(port);err!=nil{
		log.Fatal("Server not started", err)
	}

}
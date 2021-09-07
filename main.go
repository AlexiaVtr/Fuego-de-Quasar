package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func main() {

	// Variables de entorno:
	var u Url
	err := godotenv.Load()
	errEnv := envconfig.Process("split", &u)
	if err != nil || errEnv != nil {
		log.Fatal(err.Error())
	}

	// Puerto dinámico:
	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "8000"
	}
	server := NewServer(":" + Port)

	//Endpoints:
	server.Handle("/", "GET", HandleRoot)
	server.Handle(u.Host[0], "POST", HandleTopSecret)
	server.Handle(u.Host[1], "POST", SaveSatellites)
	server.Handle(u.Host[1], "GET", HandleTopSecretSplit)
	server.Handle((u.Host[1] + u.Name[0]), "POST", SaveSatellites)
	server.Handle((u.Host[1] + u.Name[1]), "POST", SaveSatellites)
	server.Handle((u.Host[1] + u.Name[2]), "POST", SaveSatellites)
	server.Handle((u.Host[1] + u.Name[3]), "GET", GetSatellites)
	server.Handle((u.Host[1] + u.Name[3]), "DELETE", DeleteSatelites)
	fmt.Println((u.Host[1] + u.Name[3]))

	server.Listen()

}

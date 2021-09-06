package main

import "os"

func main() {

	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "8000"
	}
	server := NewServer(":" + Port)

	//Endpoints:
	server.Handle("/", "GET", HandleRoot)
	server.Handle("/topsecret/", "POST", HandleTopSecret)
	server.Handle(("/topsecret_split/"), "POST", HandleTopSecret)
	server.Handle(("/topsecret_split/"), "GET", HandleTopSecretSplit)
	server.Handle("/topsecret_split/showme", "GET", GetSatellites)
	server.Handle("/topsecret_split/kenobi", "POST", SaveSatellites)
	server.Handle("/topsecret_split/skywalker", "POST", SaveSatellites)
	server.Handle("/topsecret_split/sato", "POST", SaveSatellites)

	server.Listen()
}

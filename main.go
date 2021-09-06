package main

func main() {

	//Define puerto para la función NewServe en server.go:
	server := NewServer(":8080")

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

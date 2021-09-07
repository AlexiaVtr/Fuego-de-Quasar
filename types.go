package main

type Url struct {
	Name []string
	Host []string
}

type Message struct {
	Satellites []struct {
		Name     string   `json:"name"`
		Distance float32  `json:"distance"`
		Message  []string `json:"message"`
	} `json:"satellites"`
}

type Coordenada struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type MetaData interface {
}

type Satellites []struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}
type Satellite struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

type ResponseSatellite struct {
	Position Coordenada `json:"position"`
	Message  string     `json:"message"`
}

type GetSatellite struct {
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

type AllSatellites []Satellite

var StJSON = AllSatellites{}
var S1, S2, S3 Satellite
var metadata Message
var specificMetadata []Satellite
var response ResponseSatellite
var position Coordenada
var Msg = []byte(Code200)

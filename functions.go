package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
)

// Separa la variable de la url:
func GetID(path string) string {
	var id string
	for i := 0; i != 50; i++ {
		if i >= 16 && i < len(path) {
			id += string(path[i])
		}
	}
	return id
}

func SaveSatellites(w http.ResponseWriter, r *http.Request) {
	var newSatellite Satellite
	Id := strings.TrimPrefix(r.URL.Path, "/")

	// Agregar el input:
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, Info)
	}

	// Se asigna el input como nuevo obj:
	json.Unmarshal(reqBody, &newSatellite)

	// Agrega a "Name" la variable de la url:
	newSatellite.Name = GetID(Id)

	// Respuesta en JSON:
	StJSON = append(StJSON, newSatellite)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newSatellite)
}

// Imprime lo almacenado en StJSON:
func GetSatellites(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(StJSON)
}

//Obtener localización x & y:
func GetLocation(distances ...float32) (x, y float32) {
	var Kenobi = []float32{-500, -200}
	var Skywalker = []float32{100, -100}
	var Sato = []float32{500, 100}
	d1 := distances[0]
	d2 := distances[1]
	d3 := distances[2]

	x, y = Triangulation(Kenobi[0], Kenobi[1], d1, Skywalker[0], Skywalker[1], d2, Sato[0], Sato[1], d3)

	return x, y
}

// Triangular la posición de GetLocation:
func Triangulation(x1, y1, d1, x2, y2, d2, x3, y3, d3 float32) (x, y float32) {

	A := float64(2*x2 - 2*x1)
	B := float64(2*y2 - 2*y1)
	C := math.Pow(float64(d1), 2) - math.Pow(float64(d2), 2) - math.Pow(float64(x1), 2) + math.Pow(float64(x2), 2) - math.Pow(float64(y1), 2) + math.Pow(float64(y2), 2)
	D := float64(2*x3 - 2*x2)
	E := float64(2*y3 - 2*y2)
	F := math.Pow(float64(d2), 2) - math.Pow(float64(d3), 2) - math.Pow(float64(x2), 2) + math.Pow(float64(x3), 2) - math.Pow(float64(y2), 2) + math.Pow(float64(y3), 2)
	x = float32((C*E - F*B) / (E*A - B*D))
	y = float32((C*D - A*F) / (B*D - A*E))
	return x, y
}

func GetMessage(messages ...[]string) (msg string) {

	var Kenobi []string
	var Skywalker []string
	var Sato []string
	var fixMsg []string

	// Valida que "message" tenga todos los satélites:
	if len(messages) == 3 {

		// Extrae el mensaje:
		Kenobi = append(Kenobi, messages[0]...)
		Skywalker = append(Skywalker, messages[1]...)
		Sato = append(Sato, messages[2]...)

		// Lo deja dentro de un solo slice si no es nil:
		for i := 0; i != 5; i++ {
			if Kenobi[i] != "" {
				fixMsg = append(fixMsg, Kenobi[i])

			}
			if Skywalker[i] != "" {
				fixMsg = append(fixMsg, Skywalker[i])
			}
			if Sato[i] != "" {
				fixMsg = append(fixMsg, Sato[i])
			}
		}

	}

	// Transforma el slice en string:
	message := strings.Join(fixMsg, ",")

	// Le quita las comas al string:
	messageResponse := strings.Replace(message, ",", " ", 100)
	return messageResponse
}

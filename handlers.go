package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Declaración de variables para los responses:

var Code404 string = "RESPONSE CODE: 404"
var Code200 string = "RESPONSE CODE: 200"
var Info string = "ERROR: Se necesita más información."

// Handlers de redirección:

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusExpectationFailed)
	fmt.Fprintf(w, Code404)
}

func HandleTopSecret(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&metadata)

	// Fragmentar struct de arrays para poder usar sus objetos:
	for i := range metadata.Satellites {
		specificMetadata = append(specificMetadata, metadata.Satellites[i])
	}

	if len(specificMetadata) == 3 {

		// Se agrega el mensaje con intervalo de 1 por c/ var:
		S1, S2, S3 = metadata.Satellites[0], metadata.Satellites[1], metadata.Satellites[2]
		if S1.Distance > 0 && S2.Distance > 0 && S3.Distance > 0 {
			if len(S1.Message) == 5 && len(S2.Message) == 5 && len(S3.Message) == 5 {
				// Se agrega los datos al response
				response.Message = GetMessage(S1.Message, S2.Message, S3.Message)

				// Triangular x & y para las variables del struct Coordenada.
				// Necesita las distancias de los tres satélites.
				position.X, position.Y = GetLocation(S1.Distance, S2.Distance, S3.Distance)

				response.Position = position

				// Response en formato JSON:
				data, _ := json.Marshal(response)
				w.Header().Set("Content-Type", "application/json")
				w.Write(Msg)
				w.Write(data)
				specificMetadata = specificMetadata[:0]
			} else {
				fmt.Fprintf(w, Code404)
			}
		} else {
			fmt.Fprintf(w, Code404)
		}

	} else {
		fmt.Fprintf(w, Code404)
	}
	// Control de errores:
	if err != nil {
		fmt.Fprintf(w, Code404)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func HandleTopSecretSplit(w http.ResponseWriter, r *http.Request) {
	if len(StJSON) != 3 {
		fmt.Fprintf(w, Info)
	} else {
		S1, S2, S3 = StJSON[0], StJSON[1], StJSON[2]
		// Validación de datos:
		if S1.Distance > 0 && S2.Distance > 0 && S3.Distance > 0 {
			if len(S1.Message) == 5 && len(S2.Message) == 5 && len(S3.Message) == 5 {
				position.X, position.Y = GetLocation(S1.Distance, S2.Distance, S3.Distance)
				response.Message = GetMessage(S1.Message, S2.Message, S3.Message)
				response.Position = position

				// Imprime respuesta en JSON:
				data, _ := json.Marshal(response)
				w.Header().Set("Content-Type", "application/json")
				w.Write(Msg)
				w.Write(data)

				// Resetear el slice de asteroides:
				StJSON = StJSON[:0]
			} else {
				fmt.Fprintf(w, Info)
			}
		} else {
			fmt.Fprintf(w, Info)
		}

	}
}

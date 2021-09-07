package main

import "testing"

func TestGetLocation(t *testing.T) {

	// Siempre que se envíen 3 slices en float32:
	x, y := GetLocation(1, 2, 3)
	if x == 0 || y == 0 {
		t.Error("Se recibieron coordinadas nulas")
	} else {
		t.Log("Finalizado correctamente")
	}

}

func TestGetMessage(t *testing.T) {
	S1.Message = append(S1.Message, "Prueba1", "", "", "Prueba4", "")
	S2.Message = append(S2.Message, "", "Prueba2", "", "", "Prueba5")
	S3.Message = append(S3.Message, "", "", "Prueba3", "", "")

	//Siempre que se envíen 3 slices de 5 palabras c/u:
	mensaje := GetMessage(S1.Message, S2.Message, S3.Message)
	if mensaje != "Prueba1 Prueba2 Prueba3 Prueba4 Prueba5" {
		t.Error("Falla en la unión")
	} else {
		t.Log("Finalizado correctamente")
	}

}

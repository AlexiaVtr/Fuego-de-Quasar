## Fuego-de-Quasar</br></br>
---

 <h4>⭐️ 1. Introducción</h4>

ㅤ➢ Cloud Hosting: Heroku</br>
ㅤ➢ Host: https://fuego-de-quasar-desafio.herokuapp.com</br>
ㅤ➢ Port HTTP: 8000</br>
ㅤ➢ Port HTTPS: {****}</br>

---

<h4>⭐️ 2. Lista de Operaciones</h4>
La API resuelve las siguientes consultas:</br></br>
• Operación: /topsecret/ </br>
• OperationCode: 0 </br>
• Funcionalidad: </br>
ㅤㅤ- Triangulación de una posición con coordenadas “X” e “Y” al enviarle tres distancias.</br>
ㅤㅤ- -	Recompone un mensaje troceado en tres arrays distintas.</br>
• Consideraciones: </br>
ㅤㅤ- La triangulación se realiza considerando que las distancias enviadas en el input corresponden a las siguientes coordenadas: </br>
ㅤㅤ[-500,-200], [100,-100], [500, 100]. </br>
ㅤㅤ- El mensaje debe ser de un mínimo y máximo de 5 palabras, que deben estar divididas con una iteración de una palabra por cada array:</br>
ㅤㅤEJ: [1][][][4][]</br>
ㅤㅤㅤㅤ[][2][][][5]</br>
ㅤㅤㅤㅤ[][][3][][]</br>
 </br>
• Método: POST</br>
• URI: fuego-de-quasar-desafio.herokuapp.com/topsecret/</br></br>

<h4>Request</h4>

```json
{
"satellites": [
{
"name": "kenobi",
"distance": 485.7,
"message": ["este", "", "", "mensaje", ""]
},
{
"name": "skywalker",
"distance":  266.1,
"message": ["", "es", "", "", "secreto"]
},
{
"name": "sato",
"distance": 600.5,
"message": ["", "", "un", "", ""]
}
]
}

```
<h4>Response OK</h4>

*RESPONSE CODE: 200

```json
{
 "position": 
{
 "x": -100.0,
 "y": 75.5
}, 
"message": "este es un mensaje secreto" 
}

```
<h4>Response ERROR</h4>

*RESPONSE CODE: 404</br>

---

</br>
• Operación: /topsecret_split/ </br>
• OperationCode: 1 </br>
• Funcionalidad: </br>
ㅤㅤ- Realiza las mismas funcionalidades que /topsecret/ sin un input.</br>
• Consideraciones: </br>
ㅤㅤ- Previamente se debieron haber ejecutado tres métodos POST de los servicios:</br>
ㅤㅤ“/topsecret_split/kenobi”, “/topsecret_split/skywalker”, “/topsecret_split/sato”.</br>
ㅤㅤ- La información enviada en los servicios mencionados debe corresponder con lo necesitado para las funcionalidades de la operación.</br>
ㅤㅤ- Al realizar una consulta exitosa, elimina las cadenas de JSON almacenadas.</br>
• Método: GET</br>
• URI: fuego-de-quasar-desafio.herokuapp.com/topsecret_split/</br></br>

<h4>Response OK</h4>

*RESPONSE CODE: 200

```json
{
 "position": 
{
 "x": -100.0,
 "y": 75.5
}, 
"message": "este es un mensaje secreto" 
}

```

<h4>Response ERROR</h4>

ERROR: Se necesita más información. </br>

---

</br>
• Operación: /topsecret_split/ </br>
• OperationCode: 1 </br>
• Funcionalidad: </br>
ㅤㅤ- Almacena el request JSON enviado para luego ser utilizado en el método GET de la misma operación.</br>
• Consideraciones: </br>
ㅤㅤ- No hay límite de solicitudes por operación.</br>
ㅤㅤ- De haber más de tres JSON almacenados no se podrá realizar la consulta GET de /topsecret_split/ ya que solo necesita de tres distancias/mensajes.</br>
• Método: POST</br>
• URI: fuego-de-quasar-desafio.herokuapp.com/topsecret_split/</br></br>

<h4>Request</h4>

```json
{
"name": "satellite_name",
"distance": 100.0,
"message": ["este", "", "", "mensaje", ""]
}

```

<h4>Response OK</h4>

*RESPONSE CODE: 200</br>

<h4>Response ERROR</h4>

*RESPONSE CODE: 404</br>

---

</br>

• Operación: 
ㅤㅤ-	/topsecret_split/kenobi </br>
ㅤㅤ-	/topsecret_split/skywalker </br>
ㅤㅤ-	/topsecret_split/sato </br>
 </br>
• OperationCode: 2.0, 2.1, 2.2 </br>
• Funcionalidad: </br>
ㅤㅤ- Almacena el input JSON enviado sus datos junto con el nombre del satélite (tomado de la URI), para luego ser usado en otras operaciones.</br>
• Consideraciones: </br>
ㅤㅤ- No hay límite de solicitudes por operación.</br>
ㅤㅤ- De haber más de tres JSON almacenados no se podrá realizar la consulta GET de /topsecret_split/ ya que solo necesita de tres distancias/mensajes.</br>
• Método: POST</br>
• URI: 
ㅤㅤ- fuego-de-quasar-desafio.herokuapp.com/topsecret_split/kenobi</br>
ㅤㅤ- fuego-de-quasar-desafio.herokuapp.com/topsecret_split/skywalker</br>
ㅤㅤ- fuego-de-quasar-desafio.herokuapp.com/topsecret_split/sato</br>

<h4>Request</h4>

```json
{
"distance": 100.0,
"message": ["este", "", "", "mensaje", ""]
}

```

<h4>Response OK</h4>

*RESPONSE CODE: 200</br>

<h4>Response ERROR</h4>

*RESPONSE CODE: 404</br>

---

</br>

• Operación: /topsecret_split/showme  </br>
• OperationCode: 3 </br>
• Funcionalidad: </br>
ㅤㅤ- Muestra la cadena de JSON almacenada.</br>
• Método: GET</br>
• URI: 
ㅤㅤ- fuego-de-quasar-desafio.herokuapp.com/topsecret_split/showme</br>

<h4>Response OK</h4>

*RESPONSE CODE: 200</br>

<h4>Response ERROR</h4>

*RESPONSE CODE: 404</br>

---

</br>

• Operación: /topsecret_split/showme  </br>
• OperationCode: 4 </br>
• Funcionalidad: </br>
ㅤㅤ- Elimina la cadena de JSON almacenada.</br>
• Método: DELETE</br>
• URI: 
ㅤㅤ- fuego-de-quasar-desafio.herokuapp.com/topsecret_split/showme</br>

<h4>Response OK</h4>

*RESPONSE CODE: 200</br>

<h4>Response ERROR</h4>

ERROR: No hay elementos a eliminar.</br>

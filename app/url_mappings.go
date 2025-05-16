package app

import (
	"proyecto-integrador/controllers/actividad"
	"proyecto-integrador/controllers/usuario"
)

/*
INICIO DE SESION/CREACION DE USUARIO:
	POST https://localhost:8080/login
		http_status: 201, 400, 401, 500
		body: {
			"username": "string",
			"password": "string"
		}
		respose: {
			"access_token": "TOKEN",
		}

	POST https://localhost:8080/signup
		http_status: 201, 400, 500
		body: {
			"nombre": "string",
			"apellido": "string",
			"email": "string",
			"telefono": "string"
			"username": "string",
			"password": "string"
		}


ADMINISTRACION DE LOS MODELOS:

	GET https://localhost:8080/actividades
		http_status: 200, 500
		response: {
			"actividades": [
				{
					"id": "int",
					"titulo": "string", ...
				}
			]
		}
	GET https://localhost:8080/actividades/buscar?{id,titulo,horario,categoria}
		http_status: 200, 500
		response: {
			"actividades": [
				{
					"id": "int",
					"titulo": "string", ...
				}
			]
		}
	GET https://localhost:8080/actividades/:id_actividad
		http_status: 200, 404, 400, 500
		response:
			{
				"id": "int",
				"titulo": "string", ...
			}
	// TODO: según la consigna para consultar actividades de un socio no se requiere autenticación, preguntar si está bien (sería muy raro)
	GET https://localhost:8080/usuarios/actividades
		http_status: 200, 401, 500
		header: autorization:bearer TOKEN
		response: {
			"actividades": [
				{
					"id": "int",
					"titulo": "string", ...
				}
			]
		}
	POST https://localhost:8080/inscripciones/:actividad_id
		status: 201, 404, 401, 500
		header: autorization:bearer TOKEN
		body: {
			"estado_inscripcion": "string"
		}
	DELETE https://localhost:8080/inscripciones/:actividad_id  // Permitir que un usuario elimine su inscripción en una actividad
		http_status: 204, 404, 401, 500
		header: autorization:bearer TOKEN
		body: {
			"usuario_id": "int"
		}


ENDPOINTS PARA EL ADMINISTRADOR:

	Crear una actividad
	POST https://localhost:8080/actividades (Admin)
		http_status: 201, 400, 401, 403, 500
		header: autorization:bearer TOKEN
		body: {
			"titulo": "string",
			"descripcion": "string",
			"cupo": "int",
			"dia": "string",
			"horario_inicio": "timestamp",
			"horario_final": "timestamp",
			"instructor_id": "int",
			"categoria": "string"
		}

	Actualizar una actividad
	PUT https://localhost:8080/actividades/:id (Admin)
		http_status: 200, 400, 401, 403, 500
		header: autorization:bearer TOKEN
		body: {
			"titulo": "string",
			"descripcion": "string",
			"cupo": "int",
		}
		response: <BODY>

	Borrar una actividad
	DELETE https://localhost:8080/actividades/:id (Admin) //TODO: preguntar porque al eliminar actividad, la inscripcion queda, ¿como se saca?
		http_status: 204, 404, 401, 403, 500 // TODO: no se si iba 404 con el metodo DELETE
		header: autorization:bearer TOKEN
*/

func MapURLs() {
	// TODO: capaz se puede hacer con un solo endpoint, y filtrar según si se pasan o no parametros
	router.GET("/actividades", actividad.GetAllActividades)
	router.GET("/actividades/:id", actividad.GetActividadById)
	router.GET("/actividades/buscar", actividad.GetActividadesByParams)

	router.POST("/login", usuario.Login)
}

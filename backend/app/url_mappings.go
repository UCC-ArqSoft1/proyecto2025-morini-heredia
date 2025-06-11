package app

import (
	"proyecto-integrador/controllers/actividad"
	"proyecto-integrador/controllers/inscripcion"
	"proyecto-integrador/controllers/usuario"
)

/*
INICIO DE SESION/CREACION DE USUARIO:
	POST https://localhost:8080/login
		status: 201, 400, 401, 500
		body: {
			"username": "string",
			"password": "string"
		}
		respose: {
			"access_token": "TOKEN",
		}

	POST https://localhost:8080/signup
		status: 201, 400, 500
		body: {
			"nombre": "string",
			"apellido": "string",
			"email": "string",
			"telefono": "string"
			"username": "string",
			"password": "string"
		}


ADMINISTRACION DE LOS MODELOS:

	ACTIVIDADES;
		GET https://localhost:8080/actividades
			status: 200, 400, 500
			response: {
				"actividades": [
					{
						"id": "int",
						"titulo": "string", ...
					}
				]
			}
		GET https://localhost:8080/actividades/buscar?{id,titulo,horario,categoria}
			status: 200, 500
			response: {
				"actividades": [
					{
						"id": "int",
						"titulo": "string", ...
					}
				]
			}
		GET https://localhost:8080/actividades/:id_actividad
			status: 200, 404, 400, 500
			response:
				{
					"id": "int",
					"titulo": "string", ...
				}

	USUARIOS:
		GET https://localhost:8080/usuarios/actividades
			status: 200, 401, 500
			header: autorization:bearer TOKEN
			response: {
				"actividades": [
					{
						"id": "int",
						"titulo": "string", ...
					}
				]
			}

	INCSRIPCIONES:
		GET https://localhost:8080/inscripciones
			status: 200, 401, 500
			header: autorization:bearer TOKEN
			response: {
				"inscripciones": [
					{
						"id": int,
						"fecha_inscripcion": string,
						"estado_inscripcion": string,
						...
					}, ...
				]
			}
		POST https://localhost:8080/inscripciones
			status: 201, 401, 404, 500
			header: autorization:bearer TOKEN
			body: {
				"id": int					// id de la actividad
			}
			response: {
				"id": int					// id de la inscripcion
			}

		DELETE https://localhost:8080/inscripciones  // Permitir que un usuario elimine su inscripción en una actividad
			status: 204, 404, 401, 500
			header: autorization:bearer TOKEN
			body: {
				"id": int,
				"estado_inscripcion": string
			}


ENDPOINTS PARA EL ADMINISTRADOR:

	Crear una actividad
	POST https://localhost:8080/actividades (Admin)
		status: 201, 400, 401, 403, 500
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
		status: 200, 400, 401, 403, 500
		header: autorization:bearer TOKEN
		body: {
			"titulo": "string",
			"descripcion": "string",
			"cupo": "int",
		}
		response: <BODY>

	Borrar una actividad
	DELETE https://localhost:8080/actividades/:id (Admin)
		status: 204, 404, 401, 403, 500
		header: autorization:bearer TOKEN
*/

func MapURLs() {
	// actividades
	router.GET("/actividades", actividad.GetAllActividades)
	router.GET("/actividades/:id", actividad.GetActividadById)
	router.GET("/actividades/buscar", actividad.GetActividadesByParams)
	router.POST("/actividades", JWTValidation, actividad.CreateActividad)
	router.PUT("/actividades/:id", JWTValidation, actividad.UpdateActividad)
	router.DELETE("/actividades/:id", JWTValidation, actividad.DeleteActividad)

	// usuarios
	router.POST("/login", usuario.Login)

	// inscripciones
	router.GET("/inscripciones", JWTValidation, inscripcion.GetAllInscripciones)
	router.POST("/inscripciones", JWTValidation, inscripcion.InscribirUsuario)
	router.DELETE("/inscripciones", JWTValidation, inscripcion.DesinscribirUsuario)
}

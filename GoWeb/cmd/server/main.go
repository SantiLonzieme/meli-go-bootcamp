package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Id       int     `json:"id"`
	Nombre   string  `json:"nombre"`
	Apellido string  `json:"apellido"`
	Email    string  `json:"email"`
	Edad     int     `json:"edad"`
	Altura   float64 `json:"altura"`
	Activo   bool    `json:"activo"`
	Fecha    string  `json:"fecha"`
}

type request struct {
	Id       int     `json:"id" `
	Nombre   string  `json:"nombre" binding:"required"`
	Apellido string  `json:"apellido" binding:"required"`
	Email    string  `json:"email" binding:"required"`
	Edad     int     `json:"edad" binding:"required"`
	Altura   float64 `json:"altura"`
	Activo   bool    `json:"activo"`
	Fecha    string  `json:"fecha"`
}

var usuarios []request

func main() {

	router := gin.Default()

	router.GET("/cliente/:nombre", func(ctx *gin.Context) {
		nombre := ctx.Param("nombre")
		msg := fmt.Sprintf("Hola %s", nombre)

		ctx.JSON(200, gin.H{
			"message": msg,
		})
	})

	grupoUsuarios := router.Group("/usuarios")
	grupoUsuarios.GET("/", HandlerUsers)
	grupoUsuarios.GET("/:id", UserId)
	grupoUsuarios.POST("/", Guardar())

	router.Run()

}

func Guardar() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var req request
		err := ctx.ShouldBindJSON(&req)
		token := ctx.GetHeader("token")

		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		if token != "123asd" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		req.Id = len(usuarios) + 1

		usuarios = append(usuarios, req)
		ctx.JSON(200, req)
	}

}

func HandlerUsers(ctx *gin.Context) {

	var usuarios []Usuario

	file, _ := os.ReadFile("./usuarios.json")

	err := json.Unmarshal(file, &usuarios)

	if err != nil {
		log.Fatal(err)
	}

	nombre := ctx.Query("nombre")

	if nombre != "" {
		var arrayUsuario []Usuario

		for _, usuario := range usuarios {

			if usuario.Nombre == nombre {
				arrayUsuario = append(arrayUsuario, usuario)
			}
		}

		if arrayUsuario == nil {
			ctx.JSON(404, gin.H{
				"message": "No hay usuarios con ese nombre",
			})
			return
		}

		ctx.JSON(200, arrayUsuario)
		return

	}

	apellido := ctx.Query("apellido")

	if apellido != "" {
		var arrayUsuario []Usuario

		for _, usuario := range usuarios {

			if usuario.Apellido == apellido {
				arrayUsuario = append(arrayUsuario, usuario)
			}
		}

		if arrayUsuario == nil {
			ctx.JSON(404, gin.H{
				"message": "No hay usuarios con ese apellido",
			})
			return
		}

		ctx.JSON(200, arrayUsuario)
		return

	}

	ctx.JSON(200, usuarios)

}

func UserId(ctx *gin.Context) {

	var usuarios []Usuario

	file, _ := os.ReadFile("./usuarios.json")

	err := json.Unmarshal(file, &usuarios)

	if err != nil {
		log.Fatal(err)
	}

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, usuario := range usuarios {

		if usuario.Id == id {
			ctx.JSON(200, usuario)
			return
		}
	}

	ctx.JSON(200, gin.H{
		"message": "Usuario no encontrado",
	})

}

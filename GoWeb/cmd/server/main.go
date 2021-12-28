package main

import (
	"fmt"
	"os"

	"github.com/SantiLonzieme/goweb/cmd/server/handler"
	"github.com/SantiLonzieme/goweb/internal/usuarios"
	"github.com/SantiLonzieme/goweb/pkg/store"
	"github.com/SantiLonzieme/goweb/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Usuarios API
// @version 1.0
// @description Esta Api maneja información de usuarios.
// @contact.name soporte API
// @license.name Apache 2.0
func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "usuarios.json")
	repo := usuarios.NewRepository(db)
	service := usuarios.NewService(repo)
	u := handler.NewUsuario(service)

	router := gin.Default()
	router.Use(newMiddleware)
	us := router.Group("/usuarios")
	us.POST("/", u.Store())
	us.GET("/", u.GetAll())
	us.PUT("/:id", u.Update())
	us.DELETE("/:id", u.Delete())
	us.PATCH("/:id", u.UpdateApellidoEdad())

	router.Run()

}

func newMiddleware(ctx *gin.Context) {

	token := ctx.Request.Header.Get("token")

	fmt.Println("El middleware funciona////////////", token)

	if token != os.Getenv("TOKEN") {
		ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "token inválido"))
	}

	ctx.Next()
}

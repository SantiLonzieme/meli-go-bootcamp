package main

import (
	"github.com/SantiLonzieme/goweb/cmd/server/handler"
	"github.com/SantiLonzieme/goweb/internal/usuarios"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := usuarios.NewRepository()
	service := usuarios.NewService(repo)
	u := handler.NewUsuario(service)

	router := gin.Default()
	usuarioUrl := router.Group("/usuarios")
	usuarioUrl.POST("/", u.Store())
	usuarioUrl.GET("/", u.GetAll())

	router.Run()

}

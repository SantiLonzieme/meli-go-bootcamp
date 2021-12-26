package main

import (
	"github.com/SantiLonzieme/GoWeb/cmd/server/handler"
	"github.com/SantiLonzieme/GoWeb/internal/usuarios"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := usuarios.NewRepository()
	service := usuarios.NewService(repo)
	u := handler.NewProduct(service)

	router := gin.Default()
	usuarioUrl := router.Group("/usuarios")
	usuarioUrl.POST("/", u.Store())
	usuarioUrl.GET("/", u.GetAll())

	router.Run()

}

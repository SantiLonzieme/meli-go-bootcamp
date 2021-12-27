package main

import (
	"github.com/SantiLonzieme/goweb/cmd/server/handler"
	"github.com/SantiLonzieme/goweb/internal/usuarios"
	"github.com/SantiLonzieme/goweb/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "../../usuarios.json")
	repo := usuarios.NewRepository(db)
	service := usuarios.NewService(repo)
	u := handler.NewUsuario(service)

	router := gin.Default()
	us := router.Group("/usuarios")
	us.POST("/", u.Store())
	us.GET("/", u.GetAll())
	us.PUT("/:id", u.Update())
	us.DELETE("/:id", u.Delete())
	us.PATCH("/:id", u.UpdateApellidoEdad())

	router.Run()

}

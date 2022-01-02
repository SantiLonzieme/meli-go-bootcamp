package main

import (
	"os"

	"github.com/SantiLonzieme/goweb/cmd/server/handler"
	"github.com/SantiLonzieme/goweb/docs"
	"github.com/SantiLonzieme/goweb/internal/usuarios"
	"github.com/SantiLonzieme/goweb/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones
// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	_ = godotenv.Load()

	db := store.New(store.FileType, "usuarios.json")
	repo := usuarios.NewRepository(db)
	service := usuarios.NewService(repo)
	u := handler.NewUsuario(service)

	router := gin.Default()

	us := router.Group("/usuarios")

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	us.Use(handler.NewMiddleware)
	us.POST("/", u.Store())
	us.GET("/", u.GetAll())
	us.PUT("/:id", u.Update())
	us.DELETE("/:id", u.Delete())
	us.PATCH("/:id", u.UpdateApellidoEdad())

	router.Run()
}

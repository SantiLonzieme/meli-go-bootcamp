package routes

import (
	"database/sql"

	"github.com/SantiLonzieme/sql/cmd/server/handler"
	"github.com/SantiLonzieme/sql/internal/product"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {

	repo := product.NewRepository(r.db)
	service := product.NewService(repo)
	handler := handler.NewProduct(service)

	pGroup := r.r.Group("/products")
	{
		pGroup.GET("/", handler.GetAll())
		pGroup.GET("/id/:id", handler.GetById())
		pGroup.GET("/:name", handler.GetByName())
		pGroup.POST("/", handler.Create())
		pGroup.PUT("/:id", handler.Update())
		pGroup.PUT("eficcient/:id", handler.UpdateWithContext())

	}
}

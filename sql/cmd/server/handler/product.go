package handler

import (
	"strconv"

	"github.com/SantiLonzieme/sql/internal/product"
	"github.com/SantiLonzieme/sql/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name  string  `json:"name"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

type Product struct {
	productService product.Service
}

func NewProduct(s product.Service) *Product {
	return &Product{
		productService: s,
	}
}

func (s *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		sellers, err := s.productService.GetAll()

		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}

		web.Success(ctx, 200, sellers)
	}
}

func (s *Product) GetByName() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		name := ctx.Param("name")

		product, err := s.productService.GetByName(name)

		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}

		web.Success(ctx, 200, product)
	}
}

func (s *Product) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}

		product, err := s.productService.GetById(id)

		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}

		web.Success(ctx, 200, product)
	}
}

func (s *Product) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req request

		if err := ctx.Bind(&req); err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}

		product, err := s.productService.Create(req.Name, req.Type, req.Count, req.Price)

		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}

		web.Success(ctx, 201, product)

	}
}

func (s *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}

		var req request

		if err := ctx.Bind(&req); err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}
		s, err := s.productService.Update(int(id), req.Name, req.Type, req.Count, req.Price)

		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}

		web.Success(ctx, 200, s)

	}
}

func (s *Product) UpdateWithContext() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {

			web.Error(ctx, 400, err.Error())
			return
		}

		var req request

		if err := ctx.Bind(&req); err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}
		s, err := s.productService.UpdateWithContext(ctx, int(id), req.Name, req.Type, req.Count, req.Price)

		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}

		web.Success(ctx, 200, s)

	}
}

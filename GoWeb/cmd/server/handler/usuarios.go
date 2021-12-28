package handler

import (
	"os"
	"strconv"

	"github.com/SantiLonzieme/goweb/internal/usuarios"
	"github.com/SantiLonzieme/goweb/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Id       int    `json:"id" `
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Email    string `json:"email"`
	Edad     int    `json:"edad"`
	Altura   int    `json:"altura"`
	Activo   bool   `json:"activo"`
	Fecha    string `json:"fecha"`
}

type Usuario struct {
	service usuarios.Service
}

func NewUsuario(u usuarios.Service) *Usuario {
	return &Usuario{
		service: u,
	}
}

func (u *Usuario) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}

		p, err := u.service.GetAll()

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, "Request exitoso"))
	}
}

func (u *Usuario) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}
		var req request

		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		if req.Nombre == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El nombre del usuario es requerido"))
			return
		}

		if req.Apellido == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El Apellido del usuario es requerido"))
			return
		}

		if req.Email == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El Email del usuario es requerido"))
			return
		}

		if req.Edad == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "La edad del usuario es requerido"))
			return
		}

		if req.Altura == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "La altura del usuario es requerido"))
			return
		}

		if !req.Activo {
			ctx.JSON(400, web.NewResponse(400, nil, "La propiedad activo es requerida"))
			return
		}

		if req.Fecha == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "La fecha del usuario es requerido"))
			return
		}

		p, err := u.service.Store(req.Id, req.Nombre, req.Apellido, req.Email,
			req.Edad, req.Altura, req.Activo, req.Fecha)

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, "Request exitoso"))
	}
}

func (us *Usuario) Update() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, "id inválido"))
			return
		}

		var req request

		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		if req.Nombre == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El nombre del usuario es requerido"))
			return
		}

		if req.Apellido == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El Apellido del usuario es requerido"))
			return
		}

		if req.Email == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El Email del usuario es requerido"))
			return
		}

		if req.Edad == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "La edad del usuario es requerido"))
			return
		}

		if req.Altura == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "La altura del usuario es requerido"))
			return
		}

		if !req.Activo {
			ctx.JSON(400, web.NewResponse(400, nil, "La propiedad activo es requerida"))
			return
		}

		if req.Fecha == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "La fecha del usuario es requerido"))
			return
		}

		u, err := us.service.Update(int(id), req.Nombre, req.Apellido, req.Email, req.Edad,
			req.Altura, req.Activo, req.Fecha)

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, "Request exitoso"))
	}
}

func (c *Usuario) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, "id inválido"))
			return
		}

		err = c.service.Delete(int(id))

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, id, "El usuario ha sido eliminado"))

	}
}

func (us *Usuario) UpdateApellidoEdad() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, "id inválido"))
			return
		}

		var req request

		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		if req.Apellido == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El Apellido del usuario es requerido"))
			return
		}

		if req.Edad == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "La edad del usuario es requerido"))
			return
		}

		u, err := us.service.UpdateApellidoEdad(int(id), req.Apellido, req.Edad)

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, u, "Request exitoso"))
	}
}

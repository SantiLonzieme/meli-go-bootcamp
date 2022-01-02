package handler

import (
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

// ListProducts godoc
// @Summary GetAll usuarios
// @Tags Usuario
// @Description get usuarios
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /usuarios [get]
func (u *Usuario) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// token := ctx.Request.Header.Get("token")
		// if token != os.Getenv("TOKEN") {
		// 	ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
		// 	return
		// }

		p, err := u.service.GetAll()

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, "Request exitoso"))
	}
}

// StoreProducts godoc
// @Summary Store usuarios
// @Tags Usuario
// @Description store usuarios
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param usuario body request true "Usuario to store"
// @Success 200 {object} web.Response
// @Router /usuarios [post]
func (u *Usuario) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// token := ctx.Request.Header.Get("token")
		// if token != os.Getenv("TOKEN") {
		// 	ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
		// 	return
		// }
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

// UpdateUsers godoc
// @Summary Update usuarios
// @Tags Usuario
// @Description Update usuario
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param user body request true "Usuario to update"
// @Param string query string true "Usuario ID to Update"
// @Success 200 {object} web.Response
// @Router /usuarios/{id} [put]
func (us *Usuario) Update() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// token := ctx.Request.Header.Get("token")

		// if token != os.Getenv("TOKEN") {
		// 	ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
		// 	return
		// }

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

// Delete godoc
// @Summary Delete usuarios
// @Tags Usuario
// @Description delete usuario by id
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param string query string true "Usuario ID to delete"
// @Success 200 {object} web.Response
// @Router /usuarios/{id} [delete]
func (c *Usuario) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// token := ctx.Request.Header.Get("token")

		// if token != os.Getenv("TOKEN") {
		// 	ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
		// 	return
		// }

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

// UpdateName godoc
// @Summary UpdateApellidoEdad usuarios
// @Tags Usuario
// @Description update apellido y edad de un usuario
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param string query string true "Id de usuario"
// @Success 200 {object} web.Response
// @Router /usuarios/{id} [patch]
func (us *Usuario) UpdateApellidoEdad() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// token := ctx.Request.Header.Get("token")
		// if token != os.Getenv("TOKEN") {
		// 	ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
		// 	return
		// }

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

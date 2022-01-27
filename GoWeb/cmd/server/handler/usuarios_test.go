package handler

import (
	"bytes"
	"net/http"

	"net/http/httptest"
	"os"
	"testing"

	"github.com/SantiLonzieme/goweb/internal/usuarios"
	"github.com/SantiLonzieme/goweb/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	db := store.New(store.FileType, "../usuarios.json")
	repo := usuarios.NewRepository(db)
	service := usuarios.NewService(repo)
	p := NewUsuario(service)
	r := gin.Default()

	pr := r.Group("/usuarios")
	pr.PUT("/:id", p.Update())
	pr.DELETE("/:id", p.Delete())
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func TestUpdateUsuario(t *testing.T) {
	r := createServer()
	req, rr := createRequestTest(http.MethodPut, "/usuarios/1", `{"nombre: "Up",
	"apellido": "Perez", "email": "san@gmail.com", "edad": 30, "altura": 182, "activo": true, "fecha": "12/1/2022"}`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

}

func TestDeleteUsuario(t *testing.T) {
	r := createServer()
	req, rr := createRequestTest(http.MethodDelete, "/usuarios/3", "")

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

}

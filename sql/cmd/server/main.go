package main

import (
	"database/sql"
	"fmt"

	"github.com/SantiLonzieme/sql/cmd/server/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/storage")

	if err != nil {
		fmt.Println("No se pudo conectar a la db")
	} else {
		fmt.Println("Conexion a la db exitosa !")
	}

	r := gin.Default()

	router := routes.NewRouter(r, db)
	router.MapRoutes()

	r.Run()

}

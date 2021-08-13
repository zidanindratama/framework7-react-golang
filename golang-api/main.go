package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zidanindratama/controllers"
	"github.com/zidanindratama/models"
	// swagger embed files
)

// @title Mahasiswa API
// @version 1.0
// @description API Mahasiswa Golang
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @query.collection.format multi

func main() {
	r := gin.Default()

	// Model
	db := models.SetupModels()
	r.Use(func(c *gin.Context){
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"data":"Politeknik Elektronika Negeri Surabaya"})
	})

	r.GET("/mahasiswa", controllers.MahasiswaTampil)
	r.POST("/mahasiswa", controllers.MahasiswaTambah)
	r.GET("/mahasiswa/:nim", controllers.MahasiswaDetail)
	r.PUT("/mahasiswa/:nim", controllers.MahasiswaUbah)
	r.DELETE("/mahasiswa/:nim", controllers.MahasiswaHapus)
	
	r.Run()
}
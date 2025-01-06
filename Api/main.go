package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type perro struct {
	ID     int    `json:"id"`
	Nombre string `json:"Nombre"`
	Raza   string `json:"Raza"`
	Edad   int    `json:"Edad"`
}

var perros = []perro{
	{ID: 1, Nombre: "Fido", Raza: "Labrador", Edad: 3},
	{ID: 2, Nombre: "Luna", Raza: "Golden Retriever", Edad: 5},
	{ID: 3, Nombre: "Max", Raza: "Pastor Alemán", Edad: 2},
	{ID: 4, Nombre: "Bella", Raza: "Bulldog", Edad: 4},
}

func getPerros(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, perros)

}

func crearPerro(c *gin.Context) {

	var newPerro perro

	if error := c.BindJSON(&newPerro); error != nil {
		return
	}

	if newPerro.ID == 0 || newPerro.Nombre == "" || newPerro.Raza == "" || newPerro.Edad == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos (id, Nombre, Raza, Edad) son requeridos"})
		return
	}

	perros = append(perros, newPerro)

	c.IndentedJSON(http.StatusOK, newPerro)

}

func getPerroByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	for _, perro := range perros {
		if perro.ID == id {
			c.JSON(http.StatusOK, perro)
			return
		}
	}

}

func main() {
	// Start the server
	fmt.Println("Starting the server")

	router := gin.Default()
	router.GET("/perros", getPerros)
	router.GET("/perros/:id", getPerros)

	router.POST("/perros", crearPerro)

	router.Run("localhost:8080")
}

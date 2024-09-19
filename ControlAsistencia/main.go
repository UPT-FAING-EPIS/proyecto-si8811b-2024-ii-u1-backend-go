package main

import (
	"ControlAsistencia/config"
	"ControlAsistencia/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Conectar a la base de datos
	config.ConnectDB()

	// Inicializar el router Gin
	router := gin.Default()

	// Configurar las rutas
	routes.SetupRoutes(router)

	// Iniciar el servidor en el puerto 8080
	router.Run(":8080")
}

package routes

import (
	"ControlAsistencia/controllers"
	"ControlAsistencia/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Ruta de bienvenida
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Bienvenido al sistema de control de asistencias",
		})
	})

	// Ruta para el Login
	router.POST("/login", controllers.Login)

	// Rutas protegidas por el middleware de autenticación
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware()) // Aplicar el middleware

	// Rutas para Alumnos (solo profesores autenticados pueden acceder)
	protected.GET("/alumnos", controllers.GetAlumnos)
	protected.POST("/alumnos", controllers.CreateAlumno)
	protected.PUT("/alumnos/:id", controllers.UpdateAlumno)
	protected.DELETE("/alumnos/:id", controllers.DeleteAlumno)

	// Rutas para Profesores (solo profesores autenticados pueden acceder)
	protected.GET("/profesores", controllers.GetProfesores)
	protected.POST("/profesores", controllers.CreateProfesor) // Vuelve a proteger esta ruta
	protected.PUT("/profesores/:id", controllers.UpdateProfesor)
	protected.DELETE("/profesores/:id", controllers.DeleteProfesor)

	// Rutas para Clases (solo profesores autenticados pueden acceder)
	protected.GET("/clases", controllers.GetClases)
	protected.POST("/clases", controllers.CreateClase)
	protected.PUT("/clases/:id", controllers.UpdateClase)
	protected.DELETE("/clases/:id", controllers.DeleteClase)

	// Rutas para Asistencia (solo profesores autenticados pueden acceder)
	protected.GET("/asistencias", controllers.GetAsistencias)
	protected.POST("/asistencias", controllers.CreateAsistencia)
	protected.PUT("/asistencias/:id", controllers.UpdateAsistencia)
	protected.DELETE("/asistencias/:id", controllers.DeleteAsistencia)

	// Rutas para Asignar y Desasignar Alumnos a Clases
	protected.POST("/clases/asignar-alumno", controllers.AsignarAlumnoAClase)
	protected.POST("/clases/desasignar-alumno", controllers.DesasignarAlumnoDeClase)

	// Nueva Ruta para Listar Alumnos de una Clase
	protected.GET("/clases/:id/alumnos", controllers.GetAlumnosDeClase)
}

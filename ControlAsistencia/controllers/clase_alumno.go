package controllers

import (
	"ControlAsistencia/config"
	"ControlAsistencia/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Asignar un alumno a una clase
func AsignarAlumnoAClase(c *gin.Context) {
	var input struct {
		IDClase  uint `json:"id_clase"`
		IDAlumno uint `json:"id_alumno"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de entrada no válidos", "details": err.Error()})
		return
	}

	// Crear la relación en la tabla intermedia
	claseAlumno := models.ClaseAlumno{IDClase: input.IDClase, IDAlumno: input.IDAlumno}
	if result := config.DB.Create(&claseAlumno); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al asignar el alumno a la clase", "details": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Alumno asignado a la clase correctamente"})
}

// Desasignar un alumno de una clase
func DesasignarAlumnoDeClase(c *gin.Context) {
	var input struct {
		IDClase  uint `json:"id_clase"`
		IDAlumno uint `json:"id_alumno"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de entrada no válidos", "details": err.Error()})
		return
	}

	// Eliminar la relación en la tabla intermedia
	if result := config.DB.Where("id_clase = ? AND id_alumno = ?", input.IDClase, input.IDAlumno).Delete(&models.ClaseAlumno{}); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al desasignar el alumno de la clase", "details": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Alumno desasignado de la clase correctamente"})
}

// Obtener alumnos de una clase específica
func GetAlumnosDeClase(c *gin.Context) {
	idClase := c.Param("id")
	var clase models.Clase

	// Preload para cargar los alumnos relacionados
	if err := config.DB.Preload("Alumnos").First(&clase, "id_clase = ?", idClase).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Clase no encontrada"})
		return
	}

	c.JSON(http.StatusOK, clase.Alumnos)
}

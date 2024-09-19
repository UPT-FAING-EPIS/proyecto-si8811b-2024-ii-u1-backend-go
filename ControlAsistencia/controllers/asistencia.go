package controllers

import (
	"ControlAsistencia/config"
	"ControlAsistencia/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// Obtener todas las asistencias
func GetAsistencias(c *gin.Context) {
	var asistencias []models.Asistencia
	// Usamos Preload para cargar los datos relacionados de Clase y Alumno
	if result := config.DB.Preload("Clase").Preload("Alumno").Find(&asistencias); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las asistencias"})
		return
	}
	c.JSON(http.StatusOK, asistencias)
}

// Crear una nueva asistencia
func CreateAsistencia(c *gin.Context) {
	var asistencia models.Asistencia
	if err := c.ShouldBindJSON(&asistencia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos, verifique la entrada"})
		return
	}

	// Verificar si los campos requeridos están presentes
	if asistencia.Fecha.IsZero() || asistencia.Estado == "" || asistencia.IDClase == 0 || asistencia.IDAlumno == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son obligatorios"})
		return
	}

	// Verificar que el estado sea 'presente' o 'ausente'
	if asistencia.Estado != "presente" && asistencia.Estado != "ausente" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El estado debe ser 'presente' o 'ausente'"})
		return
	}

	// Verificar si la clase y el alumno existen
	var clase models.Clase
	var alumno models.Alumno
	if err := config.DB.First(&clase, asistencia.IDClase).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Clase no encontrada"})
		return
	}
	if err := config.DB.First(&alumno, asistencia.IDAlumno).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Alumno no encontrado"})
		return
	}

	if result := config.DB.Create(&asistencia); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la asistencia"})
		return
	}

	// Cargar Clase, Profesor y Alumno después de crear la asistencia
	config.DB.Preload("Clase").Preload("Clase.Profesor").Preload("Alumno").First(&asistencia)
	c.JSON(http.StatusCreated, asistencia)
}

// Actualizar una asistencia
func UpdateAsistencia(c *gin.Context) {
	id := c.Param("id")
	var asistencia models.Asistencia

	// Verificar si la asistencia existe antes de intentar actualizar
	if result := config.DB.Preload("Clase").Preload("Alumno").First(&asistencia, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Asistencia no encontrada"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar la asistencia"})
		}
		return
	}

	if err := c.ShouldBindJSON(&asistencia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos, verifique la entrada"})
		return
	}

	// Verificar si los campos requeridos están presentes
	if asistencia.Fecha.IsZero() || asistencia.Estado == "" || asistencia.IDClase == 0 || asistencia.IDAlumno == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son obligatorios"})
		return
	}

	// Verificar que el estado sea 'presente' o 'ausente'
	if asistencia.Estado != "presente" && asistencia.Estado != "ausente" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El estado debe ser 'presente' o 'ausente'"})
		return
	}

	// Intentar guardar los cambios
	if result := config.DB.Save(&asistencia); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la asistencia"})
		return
	}

	// Cargar Clase y Alumno después de actualizar la asistencia
	config.DB.Preload("Clase").Preload("Alumno").First(&asistencia, id)
	c.JSON(http.StatusOK, asistencia)
}

// Eliminar una asistencia
func DeleteAsistencia(c *gin.Context) {
	id := c.Param("id")

	// Verificar si la asistencia existe antes de intentar eliminar
	if result := config.DB.First(&models.Asistencia{}, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Asistencia no encontrada"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar la asistencia"})
		}
		return
	}

	if result := config.DB.Delete(&models.Asistencia{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la asistencia"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Asistencia eliminada con éxito"})
}

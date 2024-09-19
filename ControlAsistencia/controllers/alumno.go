package controllers

import (
	"ControlAsistencia/config"
	"ControlAsistencia/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

// Obtener todos los alumnos
func GetAlumnos(c *gin.Context) {
	var alumnos []models.Alumno
	if result := config.DB.Find(&alumnos); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los alumnos"})
		return
	}
	c.JSON(http.StatusOK, alumnos)
}

// Crear un nuevo alumno
func CreateAlumno(c *gin.Context) {
	var alumno models.Alumno
	if err := c.ShouldBindJSON(&alumno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos, verifique la entrada"})
		return
	}

	// Verificar si los campos requeridos están presentes
	if strings.TrimSpace(alumno.Codigo) == "" || strings.TrimSpace(alumno.Nombre) == "" || strings.TrimSpace(alumno.Apellido) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son obligatorios"})
		return
	}

	if result := config.DB.Create(&alumno); result.Error != nil {
		// Verificar si el error es debido a un valor duplicado (por ejemplo, código duplicado)
		if strings.Contains(result.Error.Error(), "Duplicate entry") {
			c.JSON(http.StatusConflict, gin.H{"error": "El código del alumno ya existe"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el alumno"})
		}
		return
	}
	c.JSON(http.StatusCreated, alumno)
}

// Actualizar un alumno
func UpdateAlumno(c *gin.Context) {
	id := c.Param("id")
	var alumno models.Alumno

	// Verificar si el alumno existe antes de intentar actualizar
	if err := config.DB.First(&alumno, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Alumno no encontrado"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar el alumno"})
		}
		return
	}

	// Verificar si la entrada JSON es válida
	if err := c.ShouldBindJSON(&alumno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos, verifique la entrada"})
		return
	}

	// Verificar si los campos requeridos están presentes
	if strings.TrimSpace(alumno.Codigo) == "" || strings.TrimSpace(alumno.Nombre) == "" || strings.TrimSpace(alumno.Apellido) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son obligatorios"})
		return
	}

	// Intentar guardar los cambios
	if result := config.DB.Save(&alumno); result.Error != nil {
		if strings.Contains(result.Error.Error(), "Duplicate entry") {
			c.JSON(http.StatusConflict, gin.H{"error": "El código del alumno ya existe"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el alumno"})
		}
		return
	}
	c.JSON(http.StatusOK, alumno)
}

// Eliminar un alumno
func DeleteAlumno(c *gin.Context) {
	id := c.Param("id")
	if result := config.DB.Delete(&models.Alumno{}, id); result.Error != nil {
		// Verificar si el error es porque el registro no existe
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Alumno no encontrado"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el alumno"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Alumno eliminado con éxito"})
}

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

// Obtener todas las clases
func GetClases(c *gin.Context) {
	var clases []models.Clase
	// Utiliza Preload para cargar los datos relacionados del profesor
	if result := config.DB.Preload("Profesor").Find(&clases); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las clases"})
		return
	}
	c.JSON(http.StatusOK, clases)
}

// Crear una nueva clase
func CreateClase(c *gin.Context) {
	var clase models.Clase
	if err := c.ShouldBindJSON(&clase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos, verifique la entrada"})
		return
	}

	// Verificar si los campos requeridos están presentes
	if clase.NombreClase == "" || clase.CodigoClase == "" || clase.HorarioInicio.IsZero() || clase.HorarioFinal.IsZero() || clase.IDProfesor == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son obligatorios"})
		return
	}

	// Verificar que la clase no tenga un código duplicado
	if result := config.DB.Where("codigo_clase = ?", clase.CodigoClase).First(&models.Clase{}); result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "El código de la clase ya existe"})
		return
	}

	if result := config.DB.Create(&clase); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la clase"})
		return
	}

	// Preload para cargar el profesor después de crear la clase
	config.DB.Preload("Profesor").First(&clase)
	c.JSON(http.StatusCreated, clase)
}

// Actualizar una clase
func UpdateClase(c *gin.Context) {
	id := c.Param("id")
	var clase models.Clase

	// Verificar si la clase existe antes de intentar actualizar
	if result := config.DB.Preload("Profesor").First(&clase, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Clase no encontrada"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar la clase"})
		}
		return
	}

	// Verificar si la entrada JSON es válida
	if err := c.ShouldBindJSON(&clase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos, verifique la entrada"})
		return
	}

	// Verificar si los campos requeridos están presentes
	if clase.NombreClase == "" || clase.CodigoClase == "" || clase.HorarioInicio.IsZero() || clase.HorarioFinal.IsZero() || clase.IDProfesor == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son obligatorios"})
		return
	}

	// Intentar guardar los cambios
	if result := config.DB.Save(&clase); result.Error != nil {
		if strings.Contains(result.Error.Error(), "Duplicate entry") {
			c.JSON(http.StatusConflict, gin.H{"error": "El código de la clase ya existe"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la clase"})
		}
		return
	}

	// Preload para cargar el profesor después de actualizar la clase
	config.DB.Preload("Profesor").First(&clase, id)
	c.JSON(http.StatusOK, clase)
}

// Eliminar una clase
func DeleteClase(c *gin.Context) {
	id := c.Param("id")

	// Verificar si la clase existe antes de intentar eliminar
	if result := config.DB.First(&models.Clase{}, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Clase no encontrada"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar la clase"})
		}
		return
	}

	if result := config.DB.Delete(&models.Clase{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la clase"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Clase eliminada con éxito"})
}

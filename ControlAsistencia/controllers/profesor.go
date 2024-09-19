package controllers

import (
	"ControlAsistencia/config"
	"ControlAsistencia/models"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

// Obtener todos los profesores
func GetProfesores(c *gin.Context) {
	var profesores []models.Profesor
	if result := config.DB.Find(&profesores); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los profesores"})
		return
	}
	c.JSON(http.StatusOK, profesores)
}

// Crear un nuevo profesor
func CreateProfesor(c *gin.Context) {
	var profesor models.Profesor
	if err := c.ShouldBindJSON(&profesor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos, verifique la entrada"})
		return
	}

	// Verificar si los campos requeridos están presentes
	if strings.TrimSpace(profesor.Email) == "" || strings.TrimSpace(profesor.Contraseña) == "" || strings.TrimSpace(profesor.Nombre) == "" || strings.TrimSpace(profesor.Apellido) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son obligatorios"})
		return
	}

	// Cifrar la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(profesor.Contraseña), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al cifrar la contraseña"})
		return
	}
	profesor.Contraseña = string(hashedPassword)

	if result := config.DB.Create(&profesor); result.Error != nil {
		// Verificar si el error es debido a un valor duplicado (por ejemplo, email duplicado)
		if strings.Contains(result.Error.Error(), "Duplicate entry") {
			c.JSON(http.StatusConflict, gin.H{"error": "El email del profesor ya existe"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el profesor"})
		}
		return
	}
	c.JSON(http.StatusCreated, profesor)
}

// Actualizar un profesor
func UpdateProfesor(c *gin.Context) {
	id := c.Param("id")
	var profesor models.Profesor

	// Verificar si el profesor existe antes de intentar actualizar
	if err := config.DB.First(&profesor, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Profesor no encontrado"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar el profesor"})
		}
		return
	}

	// Verificar si la entrada JSON es válida
	if err := c.ShouldBindJSON(&profesor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos, verifique la entrada"})
		return
	}

	// Verificar si los campos requeridos están presentes
	if strings.TrimSpace(profesor.Email) == "" || strings.TrimSpace(profesor.Nombre) == "" || strings.TrimSpace(profesor.Apellido) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son obligatorios"})
		return
	}

	// Intentar guardar los cambios
	if result := config.DB.Save(&profesor); result.Error != nil {
		if strings.Contains(result.Error.Error(), "Duplicate entry") {
			c.JSON(http.StatusConflict, gin.H{"error": "El email del profesor ya existe"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el profesor"})
		}
		return
	}
	c.JSON(http.StatusOK, profesor)
}

// Eliminar un profesor
func DeleteProfesor(c *gin.Context) {
	id := c.Param("id")
	if result := config.DB.Delete(&models.Profesor{}, id); result.Error != nil {
		// Verificar si el error es porque el registro no existe
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Profesor no encontrado"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el profesor"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Profesor eliminado con éxito"})
}

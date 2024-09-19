package controllers

import (
	"ControlAsistencia/config"
	"ControlAsistencia/middleware"
	"ControlAsistencia/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// Verificar si la solicitud tiene un formato JSON válido
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de entrada no válidos"})
		return
	}

	// Validar que el email y la contraseña no estén vacíos
	if input.Email == "" || input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Correo y contraseña son obligatorios"})
		return
	}

	// Buscar profesor por email
	var profesor models.Profesor
	if err := config.DB.Where("email = ?", input.Email).First(&profesor).Error; err != nil {
		// Aquí se asume que el error es porque no se encontró el registro
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Correo o contraseña incorrectos"})
		return
	}

	// Verificar contraseña
	if err := bcrypt.CompareHashAndPassword([]byte(profesor.Contraseña), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Correo o contraseña incorrectos"})
		return
	}

	// Generar token JWT
	token, err := middleware.GenerarJWT(profesor.IDProfesor, profesor.Email)
	if err != nil {
		// Error interno al generar el token
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el token"})
		return
	}

	// Retornar el token
	c.JSON(http.StatusOK, gin.H{"token": token})
}

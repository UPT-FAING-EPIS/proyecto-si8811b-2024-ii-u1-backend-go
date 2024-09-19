package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:C25.j7e32024@tcp(127.0.0.1:3306)/control_asistencias?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	DB = database
	log.Println("Conexión a la base de datos exitosa")
}

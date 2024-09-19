package models

import (
	"time"
)

type Asistencia struct {
	ID       uint      `gorm:"column:id_asistencia;primaryKey;autoIncrement" json:"id_asistencia"`
	Fecha    time.Time `json:"fecha"`
	Estado   string    `gorm:"type:enum('presente', 'ausente');" json:"estado"`
	IDClase  uint      `gorm:"column:id_clase" json:"id_clase"`
	Clase    Clase     `gorm:"foreignKey:IDClase;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"clase"`
	IDAlumno uint      `gorm:"column:id_alumno" json:"id_alumno"`
	Alumno   Alumno    `gorm:"foreignKey:IDAlumno;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"alumno"`
}

// TableName establece el nombre de la tabla en la base de datos para el modelo Asistencia
func (Asistencia) TableName() string {
	return "asistencias"
}

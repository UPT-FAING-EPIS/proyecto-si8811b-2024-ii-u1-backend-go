package models

import (
	"time"
)

type Clase struct {
	ID            uint      `gorm:"column:id_clase;primaryKey;autoIncrement" json:"id_clase"`
	NombreClase   string    `gorm:"size:45" json:"nombre_clase"`
	CodigoClase   string    `gorm:"unique;not null;size:45" json:"codigo_clase"`
	HorarioInicio time.Time `json:"horario_inicio"`
	HorarioFinal  time.Time `json:"horario_final"`
	IDProfesor    uint      `gorm:"column:id_profesor" json:"id_profesor"`
	Profesor      Profesor  `gorm:"foreignKey:IDProfesor;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Alumnos       []Alumno  `gorm:"many2many:clase_alumnos;foreignKey:ID;joinForeignKey:IDClase;References:ID;joinReferences:IDAlumno;" json:"alumnos"`
}

// TableName establece el nombre de la tabla en la base de datos para el modelo Clase
func (Clase) TableName() string {
	return "clases"
}

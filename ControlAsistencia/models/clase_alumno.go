package models

type ClaseAlumno struct {
	ID       uint `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	IDClase  uint `gorm:"column:id_clase" json:"id_clase"`
	IDAlumno uint `gorm:"column:id_alumno" json:"id_alumno"`
}

// TableName establece el nombre de la tabla en la base de datos para el modelo ClaseAlumno
func (ClaseAlumno) TableName() string {
	return "clase_alumnos"
}

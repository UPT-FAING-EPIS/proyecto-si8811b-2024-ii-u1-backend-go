package models

type Alumno struct {
	ID       uint   `gorm:"column:id_alumno;primaryKey;autoIncrement" json:"id_alumno"`
	Codigo   string `gorm:"unique;not null" json:"codigo"`
	Nombre   string `gorm:"size:45" json:"nombre"`
	Apellido string `gorm:"size:45" json:"apellido"`
}

// TableName especifica el nombre de la tabla en la base de datos
func (Alumno) TableName() string {
	return "alumnos"
}

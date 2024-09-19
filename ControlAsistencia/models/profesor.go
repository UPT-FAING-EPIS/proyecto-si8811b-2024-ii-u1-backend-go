package models

type Profesor struct {
	IDProfesor uint   `gorm:"column:id_profesor;primaryKey;autoIncrement" json:"id_profesor"`
	Nombre     string `gorm:"size:50" json:"nombre"`
	Apellido   string `gorm:"size:45" json:"apellido"`
	Email      string `gorm:"unique;not null" json:"email"`
	Contraseña string `gorm:"size:45" json:"contraseña"`
}

// TableName especifica el nombre de la tabla en la base de datos
func (Profesor) TableName() string {
	return "profesores"
}

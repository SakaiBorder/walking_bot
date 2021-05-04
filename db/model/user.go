package model

import "database/sql"

type User struct {
	ID        uint `gorm:"primary_key"`
	LineID    string
	Longitude sql.NullFloat64
	Latitude  sql.NullFloat64
	Distance  sql.NullInt32
}

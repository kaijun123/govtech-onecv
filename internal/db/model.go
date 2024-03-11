package db

import "github.com/lib/pq"

type StudentSchema struct {
	Student  string         `json:"student" gorm:"primaryKey"`
	Suspend  bool           `json:"suspend,omitempty"`
	Teachers pq.StringArray `json:"teachers,omitempty" gorm:"type:text[]"`
}

type TeacherSchema struct {
	Teacher  string         `json:"teacher" gorm:"primaryKey"`
	Students pq.StringArray `json:"students,omitempty" gorm:"type:text[]"`
}

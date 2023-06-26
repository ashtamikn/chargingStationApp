package models

type MockTable struct {
	EndPoint string `gorm:"size:255"`
	Method   string
	Response string `gorm:"type:jsonb"`
}

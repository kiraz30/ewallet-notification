package models

import "time"

type NotificationTemplate struct {
	ID           int
	TemplateName string `gorm:"column:template_name; type:varchar(255);unique"`
	Subject      string `gorm:"column:subject; type:varchar(255)"`
	Body         string `gorm:"column:body; type:text"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (*NotificationTemplate) TableName() string {
	return "notification_templates"
}

type NotificationHistory struct {
	ID           int
	RecipientID  string `gorm:"column:recipient"`
	TemplateID   int    `gorm:"column:template_id; type:varchar(255)"`
	Status       string `gorm:"column:status; type:varchar(10)"`
	ErrorMessage string `gorm:"column:error_message; type:text"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (*NotificationHistory) TableName() string {
	return "notification_history"
}

package model

import (
	"time"

	"github.com/google/uuid"
	softDel "gorm.io/plugin/soft_delete"
)

type Subject struct {
	ID        uuid.UUID `gorm:"primaryKey; column:id; type:uuid; default:gen_random_uuid();" json:"id"`
	Name      string    `gorm:"column:name; not null; size:255;" json:"name"`
	Credits   uint8     `gorm:"column:credits; not null;" json:"credits"`
	
	Students  []Student `gorm:"many2many:student_subjects;" json:"students,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime; column:created_at;" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime; column:updated_at;" json:"-"`
	DeletedAt time.Time `gorm:"column:deleted_at;" json:"-"`
	IsDeleted softDel.DeletedAt `gorm:"column:is_deleted; softDelete:flag; DeletedAtField:DeletedAt;" json:"-"`
}

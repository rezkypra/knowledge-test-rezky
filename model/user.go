package model

import (
	"ktfs/config"

	"time"

	"golang.org/x/crypto/bcrypt"
	softDel "gorm.io/plugin/soft_delete"
	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `gorm:"primaryKey; column:id; type:uuid; default:gen_random_uuid();" json:"id"`
	Name string `gorm:"column:name; default:null; size:255;" json:"name"`
	Email string `gorm:"unique; column:email; size:255; uniqueIndex:udx_email_user;" json:"email"`
	Password string `gorm:"column:password; default:null; size:255;" json:"-"`
	Token string `gorm:"column:token; default:null; size:255;" json:"token"`

	CreatedAt time.Time `gorm:"autoCreateTime; column:created_at;" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime; column:updated_at;" json:"-"`
	DeletedAt	time.Time `gorm:"column:deleted_at;" json:"-"`
	IsDeleted softDel.DeletedAt `gorm:"column:is_deleted; softDelete:flag; DeletedAtField:DeletedAt; uniqueIndex:udx_email_user;" json:"-"`
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func IsUserExistsByEmail(email string) (bool, error) {
	var userFound int64
	err := config.DB.Model(User{}).Where("email  = ?", email).Count(&userFound).Error
	if err != nil {
		return false, err
	}
	if userFound == 0 {
		return false, nil
	} else {
		return true, nil
	}
}
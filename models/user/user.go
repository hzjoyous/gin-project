package user

import "thousand-hands-server/models"

// User 用户模型
type User struct {
	models.BaseModel

	Name     string `gorm:"type:varchar(255);not null;unique" valid:"name"`
	Email    string `gorm:"type:varchar(255);unique;" valid:"email"`
	Password string `gorm:"type:varchar(255)" valid:"password"`

	// gorm:"-" —— 设置 GORM 在读写时略过此字段，仅用于表单验证
	PasswordComfirm string `gorm:"-" valid:"password_comfirm"`

	// ArticleID uint64 `gorm:"not null;index"`
	// Article   article.Article
}

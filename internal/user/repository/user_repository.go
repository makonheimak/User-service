package repository

import (
	"github.com/makonheimak/user-service/internal/user/orm"

	"gorm.io/gorm"
)

type UserRepository interface {
	PostUser(req *orm.User) error
	GetAllUsers() ([]orm.User, error)
	GetUserByID(id int64) (orm.User, error)
	PatchUserByID(user *orm.User) error
	DeleteUserByID(id int64) error
}

type Repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) PostUser(req *orm.User) error {
	return r.db.Create(req).Error
}

func (r *Repository) GetAllUsers() ([]orm.User, error) {
	var users = []orm.User{}
	err := r.db.Find(&users).Error
	return users, err
}

func (r *Repository) GetUserByID(id int64) (orm.User, error) {
	var user orm.User
	err := r.db.First(&user, "id = ?", id).Error
	return user, err
}

func (r *Repository) PatchUserByID(user *orm.User) error { // ← изменить на указатель
	return r.db.Model(user).Updates(map[string]interface{}{
		"email":    user.Email,
		"password": user.Password,
	}).Error
}

func (r *Repository) DeleteUserByID(id int64) error {
	return r.db.Delete(&orm.User{}, "id = ?", id).Error
}

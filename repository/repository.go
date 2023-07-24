package repository

import (
	"desa-sragen/domain"
	"errors"

	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) domain.Repo {
	return &repo{
		db: db,
	}
}

func (r *repo) GetUser(username string) (record *domain.Auth, err error) {
	if err := r.db.Table("auth").Where("username = ?", username).Scan(&record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

func (r *repo) CreateUser(req domain.Auth) error {
	if err := r.db.Table("Auth").Create(&req).Error; err != nil {
		return errors.New("internal server error")
	}

	return nil
}

func (r *repo) GetAllAdmin() (record []domain.Auth, err error) {
	if err := r.db.Table("auth").Where("role_id = 2").Scan(&record).Error; err != nil {
		return nil, errors.New("internal server error")
	}

	return record, nil  
}

func (r *repo) InActiveAuth(id int, isActive bool) error {
	data := map[bool]int8{true:1, false:0}[isActive]

	if err := r.db.Debug().Table("auth").Where("auth_id = ?", id).Update("is_active", data).Error; err != nil {
		return errors.New("internal server error")
	}

	return nil
}
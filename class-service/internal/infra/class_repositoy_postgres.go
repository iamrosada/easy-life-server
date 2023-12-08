package repository

import (
	"github.com/iamrosada/easy-life-server/class-service/internal/entity"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ClassRepositoryPostgres struct {
	DB *gorm.DB
}

func NewClassRepositoryPostgres(db *gorm.DB) *ClassRepositoryPostgres {
	return &ClassRepositoryPostgres{DB: db}
}

func (r *ClassRepositoryPostgres) Create(Class *entity.Class) error {
	return r.DB.Create(Class).Error
}

func (r *ClassRepositoryPostgres) FindAll() ([]*entity.Class, error) {
	var Classs []*entity.Class
	if err := r.DB.Find(&Classs).Error; err != nil {
		return nil, err
	}
	return Classs, nil
}

func (r *ClassRepositoryPostgres) Update(Class *entity.Class) error {
	return r.DB.Save(Class).Error
}

func (r *ClassRepositoryPostgres) DeleteByID(id uint) error {
	return r.DB.Where("id = ?", id).Delete(entity.Class{}).Error
}

func (r *ClassRepositoryPostgres) GetByID(id uint) (*entity.Class, error) {
	var Class entity.Class
	if err := r.DB.Where("id = ?", id).First(&Class).Error; err != nil {
		return nil, err
	}
	return &Class, nil
}

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

func (r *ClassRepositoryPostgres) Create(class *entity.Class) error {
	return r.DB.Create(class).Error
}

func (r *ClassRepositoryPostgres) FindAll() ([]*entity.Class, error) {
	var Classs []*entity.Class
	if err := r.DB.Find(&Classs).Error; err != nil {
		return nil, err
	}
	return Classs, nil
}

func (r *ClassRepositoryPostgres) Update(class *entity.Class) error {
	return r.DB.Save(&class).Error

}

func (r *ClassRepositoryPostgres) DeleteByID(id string) error {
	return r.DB.Where("id = ?", id).Delete(entity.Class{}).Error
}

func (r *ClassRepositoryPostgres) GetByID(id string) (*entity.Class, error) {
	var class entity.Class
	if err := r.DB.Where("id = ?", id).First(&class).Error; err != nil {
		return nil, err
	}
	return &class, nil
}

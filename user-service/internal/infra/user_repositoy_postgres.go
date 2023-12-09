package repository

import (
	"github.com/iamrosada/easy-life-server/user-server/internal/entity"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type StudentRepositoryPostgres struct {
	DB *gorm.DB
}

func NewStudentRepositoryPostgres(db *gorm.DB) *StudentRepositoryPostgres {
	return &StudentRepositoryPostgres{DB: db}
}

func (r *StudentRepositoryPostgres) Create(Student *entity.Student) error {
	return r.DB.Create(Student).Error
}

func (r *StudentRepositoryPostgres) FindAll() ([]*entity.Student, error) {
	var Students []*entity.Student
	if err := r.DB.Find(&Students).Error; err != nil {
		return nil, err
	}
	return Students, nil
}

func (r *StudentRepositoryPostgres) Update(Student *entity.Student) error {
	return r.DB.Save(Student).Error
}

func (r *StudentRepositoryPostgres) DeleteByID(id string) error {
	return r.DB.Where("id = ?", id).Delete(entity.Student{}).Error
}

func (r *StudentRepositoryPostgres) GetByID(id string) (*entity.Student, error) {
	var Student entity.Student
	if err := r.DB.Where("id = ?", id).First(&Student).Error; err != nil {
		return nil, err
	}
	return &Student, nil
}

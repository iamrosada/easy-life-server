package repository

import (
	"github.com/iamrosada/easy-life-server/user-server/internal/entity"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TeacherRepositoryPostgres struct {
	DB *gorm.DB
}

func NewTeacherRepositoryPostgres(db *gorm.DB) *TeacherRepositoryPostgres {
	return &TeacherRepositoryPostgres{DB: db}
}

func (r *TeacherRepositoryPostgres) Create(Teacher *entity.Teacher) error {
	return r.DB.Create(Teacher).Error
}

func (r *TeacherRepositoryPostgres) FindAll() ([]*entity.Teacher, error) {
	var Teachers []*entity.Teacher
	if err := r.DB.Find(&Teachers).Error; err != nil {
		return nil, err
	}
	return Teachers, nil
}

func (r *TeacherRepositoryPostgres) Update(Teacher *entity.Teacher) error {
	return r.DB.Save(Teacher).Error
}

func (r *TeacherRepositoryPostgres) DeleteByID(id string) error {
	return r.DB.Where("id = ?", id).Delete(entity.Teacher{}).Error
}

func (r *TeacherRepositoryPostgres) GetByID(id string) (*entity.Teacher, error) {
	var Teacher entity.Teacher
	if err := r.DB.Where("id = ?", id).First(&Teacher).Error; err != nil {
		return nil, err
	}
	return &Teacher, nil
}

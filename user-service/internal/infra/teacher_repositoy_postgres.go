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

func (r *TeacherRepositoryPostgres) Create(teacher *entity.Teacher) error {
	return r.DB.Create(teacher).Error
}

func (r *TeacherRepositoryPostgres) FindAll() ([]*entity.Teacher, error) {
	var teachers []*entity.Teacher
	if err := r.DB.Find(&teachers).Error; err != nil {
		return nil, err
	}
	return teachers, nil
}

func (r *TeacherRepositoryPostgres) Update(teacher *entity.Teacher) error {
	return r.DB.Save(teacher).Error
}

func (r *TeacherRepositoryPostgres) DeleteByID(id string) error {
	return r.DB.Where("id = ?", id).Delete(entity.Teacher{}).Error
}

func (r *TeacherRepositoryPostgres) GetByID(id string) (*entity.Teacher, error) {
	var teacher entity.Teacher
	if err := r.DB.Where("id = ?", id).First(&teacher).Error; err != nil {
		return nil, err
	}
	return &teacher, nil
}

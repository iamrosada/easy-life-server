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

func (r *StudentRepositoryPostgres) Create(student *entity.Student) error {

	createdNewStudent := entity.Student{
		ID:          student.ID,
		TeachersIDs: student.TeachersIDs,
		FullName:    student.FullName,
		Name:        student.Name,
		CourseName:  student.CourseName,
	}

	if err := r.DB.Create(&createdNewStudent).Error; err != nil {
		return err
	}

	return nil
}

func (r *StudentRepositoryPostgres) FindAll() ([]*entity.Student, error) {
	var students []*entity.Student

	// Fetching teachers_ids for all students
	if err := r.DB.Find(&students).Error; err != nil {
		return nil, err
	}

	return students, nil
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

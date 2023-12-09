package repository

import (
	"fmt"
	"strings"

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
	// Ensure unique teacher IDs
	teacherMap := make(map[string]bool)
	for _, teacherID := range student.TeachersIDs {
		lowercaseID := strings.ToLower(teacherID)

		if _, exists := teacherMap[lowercaseID]; exists {
			return fmt.Errorf("duplicate teacher ID found: %s", teacherID)
		}

		teacherMap[lowercaseID] = true

		newStudent := entity.Student{
			ID:          student.ID, // Generate a new UUID for the student ID
			TeachersIDs: []string{teacherID},
			Name:        student.Name,
			CourseName:  student.CourseName,
		}

		// Create the student record in the database
		if err := r.DB.Create(&newStudent).Error; err != nil {
			return err
		}

	}

	// Create the student

	return nil
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

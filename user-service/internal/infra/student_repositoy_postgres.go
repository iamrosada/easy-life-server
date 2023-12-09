package repository

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
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

	fmt.Println("Debug information:", student.ID)

	studentMap := make(map[string]bool)
	for _, teacherIDsValue := range student.TeachersIDs {
		lowercaseStudent := strings.ToLower(teacherIDsValue)

		if _, exists := studentMap[lowercaseStudent]; exists {
			return fmt.Errorf("duplicate code found: %s", teacherIDsValue)
		}

		studentMap[lowercaseStudent] = true

		student := entity.Student{
			ID:          uuid.New().String(),
			StudentID:   student.ID,
			TeachersIDs: []string{teacherIDsValue},
			FullName:    student.FullName,
			Name:        student.Name,
			CourseName:  student.CourseName,
		}

		if err := r.DB.Create(&student).Error; err != nil {
			return err
		}
	}

	return nil
}

// func (r *StudentRepositoryPostgres) FindAll() ([]*entity.Student, error) {
// 	var students []*entity.Student
// 	var teachers_ids []string
// 	err := r.DB.Table("student").Select("teachers_ids").Pluck("teachers_ids", &teachers_ids).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := r.DB.Select("id", "name", "student_id", "full_name", "course_name").Find(&students).Error; err != nil {
// 		return nil, err
// 	}

//		return students, nil
//	}

func (r *StudentRepositoryPostgres) FindAll() ([]*entity.Student, error) {
	var students []*entity.Student

	// Fetching teachers_ids for all students
	var teachersIDs []string
	if err := r.DB.Table("students").Pluck("teachers_ids", &teachersIDs).Error; err != nil {
		return nil, err
	}

	// Fetching other student details
	if err := r.DB.Select("id", "name", "student_id", "full_name", "course_name").Find(&students).Error; err != nil {
		return nil, err
	}

	// Splitting the comma-separated string into a slice for each student
	for i, student := range students {
		student.TeachersIDs = strings.Split(teachersIDs[i], ",")
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

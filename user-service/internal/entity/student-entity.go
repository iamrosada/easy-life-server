package entity

import (
	"errors"

	"github.com/google/uuid"
)

type StudentRepository interface {
	Create(student *Student) error
	FindAll() ([]*Student, error)
	Update(student *Student) error
	DeleteByID(id string) error
	GetByID(id string) (*Student, error)
}

type Student struct {
	ID         string
	Name       string
	FullName   string
	CourseName string
	TeacherID  string
	// Teachers   Teacher `gorm:"foreignKey:TeacherID"`
	Teachers []Teacher `gorm:"many2many:student_teachers;"`
}

func NewStudent(name, full_name, course_name string) *Student {
	return &Student{
		ID:         uuid.New().String(),
		Name:       name,
		FullName:   full_name,
		CourseName: course_name,
	}
}

func (d *Student) Update(course_name, full_Name, name string) {
	d.CourseName = course_name
	d.FullName = full_Name
	d.Name = name
}

type InMemoryStudentRepository struct {
	Students map[string]*Student
}

func NewInMemoryStudentRepository() *InMemoryStudentRepository {
	return &InMemoryStudentRepository{
		Students: make(map[string]*Student),
	}
}

func (r *InMemoryStudentRepository) DeleteByID(id string) error {
	if _, exists := r.Students[id]; !exists {
		return errors.New("Student not found")
	}

	delete(r.Students, id)
	return nil
}

func (r *InMemoryStudentRepository) FindAll() ([]*Student, error) {
	var allStudents []*Student
	for _, Student := range r.Students {
		allStudents = append(allStudents, Student)
	}
	return allStudents, nil
}

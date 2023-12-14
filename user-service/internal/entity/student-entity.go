package entity

import (
	"errors"

	"github.com/google/uuid"
)

type StudentRepository interface {
	Create(student *Student) error
	FindAll() ([]*Student, error)
	ListStudentsByTeacherID(teacherId string) ([]*Student, error)

	Update(student *Student) error
	DeleteByID(id string) error
	ApplyEvent(eventID string, StudentdsIDs []string) error
	GetByID(id string) (*Student, error)
	GetByEmail(id string) (*Student, error)
	GetByEventID(id string) ([]*Student, error)
}
type Serializable interface {
	BeforeSave() error
	AfterFind() error
}
type TeachersIDs []string

type Student struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	FullName    string      `json:"full_name"`
	CourseName  string      `json:"course_language"`
	TeachersIDs TeachersIDs `gorm:"type:VARCHAR(255)" json:"teachers_ids"`
	EventID     string      `json:"event_id"`
	Email       string      `json:"email"`
}

func NewStudent(name, full_name, course_name, email string) *Student {
	return &Student{
		ID:         uuid.New().String(),
		Name:       name,
		FullName:   full_name,
		CourseName: course_name,
		Email:      email,
	}
}

func (d *Student) Update(course_name, full_Name, name, event_id, email string) {
	d.CourseName = course_name
	d.FullName = full_Name
	d.Name = name
	d.Email = email

	d.EventID = event_id
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

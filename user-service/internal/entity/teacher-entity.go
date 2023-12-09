package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type TeacherRepository interface {
	Create(teacher *Teacher) error
	FindAll() ([]*Teacher, error)
	Update(teacher *Teacher) error
	DeleteByID(id string) error
	GetByID(id string) (*Teacher, error)
}

type Teacher struct {
	ID             string `gorm:"primaryKey"`
	Name           string `json:"name"`
	CourseLanguage string `json:"course_language"`
	IsActive       bool   `json:"is_active"`
}

func NewTeacher(name, course_language string, is_active bool) *Teacher {
	return &Teacher{
		ID:             uuid.New().String(),
		Name:           name,
		CourseLanguage: course_language,
		IsActive:       is_active,
	}
}

func (d *Teacher) Update(name, course_language string, is_active bool) {
	d.Name = name
	d.CourseLanguage = course_language
	d.IsActive = is_active
}

type InMemoryTeacherRepository struct {
	Teachers map[string]*Teacher
}

func NewInMemoryTeacherRepository() *InMemoryTeacherRepository {
	return &InMemoryTeacherRepository{
		Teachers: make(map[string]*Teacher),
	}
}

func (r *InMemoryTeacherRepository) DeleteByID(id string) error {
	if _, exists := r.Teachers[id]; !exists {
		return errors.New("Teacher not found")
	}

	delete(r.Teachers, id)
	return nil
}

func (r *InMemoryTeacherRepository) FindAll() ([]*Teacher, error) {
	var allTeachers []*Teacher
	for _, Teacher := range r.Teachers {
		allTeachers = append(allTeachers, Teacher)
	}
	return allTeachers, nil
}

func (t Teacher) Value() (driver.Value, error) {
	return json.Marshal(t)
}

// Scan implements the sql.Scanner interface.
func (t *Teacher) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	data, ok := value.([]byte)
	if !ok {
		return errors.New("Invalid scan type for Teacher")
	}
	return json.Unmarshal(data, t)
}

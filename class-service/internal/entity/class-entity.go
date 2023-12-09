package entity

import (
	"errors"

	"github.com/google/uuid"
)

type ClassRepository interface {
	Create(class *Class) error
	FindAll() ([]*Class, error)
	Update(Class *Class) error
	DeleteByID(id string) error
	GetByID(id string) (*Class, error)
}

type Serializable interface {
	BeforeSave() error
	AfterFind() error
}
type StudentsIDs []string
type Class struct {
	ID            string      `json:"id"`
	ClassID       string      `json:"class_id"`
	TitleOfLesson string      `json:"title_of_lesson"`
	Description   string      `json:"description"`
	TeacherID     string      `json:"teacher_id"`
	StudentsIDs   StudentsIDs `gorm:"type:VARCHAR(255)" json:"students_ids"`
}

func NewClass(titleOfLesson, description, teacherID string, studentsIDs []string) *Class {
	return &Class{
		ID:            uuid.New().String(),
		TitleOfLesson: titleOfLesson,
		Description:   description,
		TeacherID:     teacherID,
		StudentsIDs:   studentsIDs,
	}
}

func (c *Class) Update(titleOfLesson, description, teacherID string) {
	c.TitleOfLesson = titleOfLesson
	c.Description = description
	c.TeacherID = teacherID
}

type InMemoryClassRepository struct {
	Classs map[string]*Class
}

func NewInMemoryClassRepository() *InMemoryClassRepository {
	return &InMemoryClassRepository{
		Classs: make(map[string]*Class),
	}
}

func (r *InMemoryClassRepository) DeleteByID(id string) error {
	if _, exists := r.Classs[id]; !exists {
		return errors.New("Class not found")
	}

	delete(r.Classs, id)
	return nil
}

func (r *InMemoryClassRepository) FindAll() ([]*Class, error) {
	var allClasss []*Class
	for _, Class := range r.Classs {
		allClasss = append(allClasss, Class)
	}
	return allClasss, nil
}

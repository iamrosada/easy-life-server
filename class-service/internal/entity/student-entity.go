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

type Class struct {
	ID         string
	Name       string
	FullName   string
	CourseName string
}

func NewClass(name, full_name, course_name string) *Class {
	return &Class{
		ID:         uuid.New().String(),
		Name:       name,
		FullName:   full_name,
		CourseName: course_name,
	}
}

func (d *Class) Update(course_name, full_Name, name string) {
	d.CourseName = course_name
	d.FullName = full_Name
	d.Name = name
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

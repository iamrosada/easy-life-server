package repository

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/iamrosada/easy-life-server/class-service/internal/entity"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ClassRepositoryPostgres struct {
	DB *gorm.DB
}

func NewClassRepositoryPostgres(db *gorm.DB) *ClassRepositoryPostgres {
	return &ClassRepositoryPostgres{DB: db}
}

func (r *ClassRepositoryPostgres) Create(class *entity.Class) error {
	// Iterate through each student ID and create a new class for each
	createdNewClass := entity.Class{
		ID:            uuid.New().String(),
		ClassID:       class.ID,
		TitleOfLesson: class.TitleOfLesson,
		Description:   class.Description,
		TeacherID:     class.TeacherID,
		StudentsIDs:   class.StudentsIDs,
	}
	// []string{"vida", "Amor"}
	if err := r.DB.Create(&createdNewClass).Error; err != nil {
		// log.Printf("Error creating class for user %s: %v", studentID, err)
		// return fmt.Errorf("failed to create class for user %s: %v", studentID, err)

	}
	return nil
}

func (r *ClassRepositoryPostgres) FindAll() ([]*entity.Class, error) {
	var class []*entity.Class
	marshaled, _ := json.MarshalIndent(class, "", "\t")
	fmt.Println(string(marshaled))
	if err := r.DB.Find(&class).Error; err != nil {
		return nil, err
	}

	return class, nil
}

func (r *ClassRepositoryPostgres) Update(class *entity.Class) error {
	return r.DB.Save(&class).Error

}

func (r *ClassRepositoryPostgres) DeleteByID(id string) error {
	return r.DB.Where("id = ?", id).Delete(entity.Class{}).Error
}

func (r *ClassRepositoryPostgres) GetByID(id string) (*entity.Class, error) {
	var class entity.Class
	if err := r.DB.Where("id = ?", id).First(&class).Error; err != nil {
		return nil, err
	}
	return &class, nil
}

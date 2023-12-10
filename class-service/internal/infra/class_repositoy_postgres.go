package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

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
		ID:            class.ID,
		TitleOfLesson: class.TitleOfLesson,
		Description:   class.Description,
		TeacherID:     class.TeacherID,
		StudentsIDs:   class.StudentsIDs,
	}

	if err := r.DB.Create(&createdNewClass).Error; err != nil {
		return fmt.Errorf("failed to create class: %v", err)
	}

	// Call the external API with event ID and student IDs
	if err := r.callExternalAPI(class.ID, class.StudentsIDs); err != nil {
		// Handle the error as needed
		return fmt.Errorf("failed to call external API: %v", err)
	}

	return nil
}

func (r *ClassRepositoryPostgres) callExternalAPI(eventID string, studentIDs []string) error {
	apiURL := fmt.Sprintf("http://localhost:8000/api/students/event/%s", eventID)

	// Create a map for the request body
	requestBody := map[string]interface{}{
		"students_ids": studentIDs,
	}

	// Convert the map to JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	// Make a POST request to the external API
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to make POST request to external API: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response status from external API: %v", resp.Status)
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

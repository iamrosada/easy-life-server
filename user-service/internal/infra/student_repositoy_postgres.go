package repository

import (
	"fmt"
	"log"

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
		Email:       student.Email,
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

func (r *StudentRepositoryPostgres) GetByEventID(id string) ([]*entity.Student, error) {
	var students []*entity.Student

	// Fetching teachers_ids for all students
	if err := r.DB.Where("event_id = ?", id).Find(&students).Error; err != nil {
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

func (r *StudentRepositoryPostgres) GetByEmail(email string) (*entity.Student, error) {
	var Student entity.Student
	if err := r.DB.Where("email = ?", email).First(&Student).Error; err != nil {
		return nil, err
	}
	return &Student, nil
}
func (r *StudentRepositoryPostgres) ApplyEvent(eventID string, studentIDs []string) error {
	// Iterate through each student ID and apply the event
	for _, studentID := range studentIDs {
		// Fetch student information using GetByID
		getStudent, err := r.GetByID(studentID)
		if err != nil {
			log.Printf("Error fetching student information for ID %s: %v", studentID, err)
			return fmt.Errorf("failed to fetch student information for ID %s: %v", studentID, err)
		}

		// Update the EventID field for the event being applied
		getStudent.EventID = eventID

		log.Printf("Applying event for student %s with event ID %s", studentID, eventID)

		// Save the updated student record
		if err := r.DB.Save(getStudent).Error; err != nil {
			log.Printf("Error applying event for student %s: %v", studentID, err)
			return fmt.Errorf("failed to apply event for student %s: %v", studentID, err)
		}
	}

	return nil
}

func (r *StudentRepositoryPostgres) ListStudentsByTeacherID(teacherID string) ([]*entity.Student, error) {
	var students []*entity.Student

	// Fetching all students
	if err := r.DB.Find(&students).Error; err != nil {
		return nil, err
	}

	// Filter students by teacher ID
	filteredStudents := make([]*entity.Student, 0)
	for _, student := range students {
		for _, teacherIDInList := range student.TeachersIDs {
			if teacherIDInList == teacherID {
				// Add the student to the filtered list
				filteredStudents = append(filteredStudents, student)
				break // No need to check further once a match is found
			}
		}
	}

	return filteredStudents, nil
}

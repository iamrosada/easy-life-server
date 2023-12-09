package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/iamrosada/easy-life-server/class-service/internal/entity"
)

type CreateClassInputDto struct {
	TitleOfLesson string   `json:"title_of_lesson"`
	Description   string   `json:"description"`
	TeacherID     string   `json:"teacher_id"`
	StudentsIDs   []string `json:"students_ids"`
}

type CreateClassOutputDto struct {
	ClassID       string   `json:"class_id"`
	TitleOfLesson string   `json:"title_of_lesson"`
	Description   string   `json:"description"`
	TeacherID     string   `json:"teacher_id"`
	StudentsIDs   []string `json:"students_ids"`
}

type CreateClassUseCase struct {
	ClassRepository entity.ClassRepository
}

func NewCreateClassUseCase(ClassRepository entity.ClassRepository) *CreateClassUseCase {
	return &CreateClassUseCase{ClassRepository: ClassRepository}
}

func (u *CreateClassUseCase) Execute(input CreateClassInputDto) (*CreateClassOutputDto, error) {
	// Validate input if necessary
	// Create a new class entity
	newClass := entity.NewClass(input.TitleOfLesson, input.Description, input.TeacherID, input.StudentsIDs)
	newClass.StudentsIDs = input.StudentsIDs

	marshaled, _ := json.MarshalIndent(newClass, "", "\t")
	fmt.Println(string(marshaled))
	// Save the new class using the repository
	if err := u.ClassRepository.Create(newClass); err != nil {
		return nil, err
	}

	// Return the output DTO with relevant information
	outputDto := &CreateClassOutputDto{
		ClassID:       newClass.ID,
		TitleOfLesson: newClass.TitleOfLesson,
		Description:   newClass.Description,
		TeacherID:     newClass.TeacherID,
		StudentsIDs:   newClass.StudentsIDs,
	}

	return outputDto, nil
}

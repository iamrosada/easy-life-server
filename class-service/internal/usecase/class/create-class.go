package usecase

import (
	"github.com/iamrosada/easy-life-server/class-service/internal/entity"
)

type CreateClassInputDto struct {
	TitleOfLesson string   `json:"title_of_lesson"`
	Description   string   `json:"description"`
	TeacherID     string   `json:"teacher_id"`
	StudentsIDs   []string `json:"students_ids"`
	GoogleMeetUrl string   `json:"google_meet_url"`
}

type CreateClassOutputDto struct {
	ID            string   `json:"id"`
	TitleOfLesson string   `json:"title_of_lesson"`
	Description   string   `json:"description"`
	TeacherID     string   `json:"teacher_id"`
	StudentsIDs   []string `json:"students_ids"`
	GoogleMeetUrl string   `json:"google_meet_url"`
}

type CreateClassUseCase struct {
	ClassRepository entity.ClassRepository
}

func NewCreateClassUseCase(ClassRepository entity.ClassRepository) *CreateClassUseCase {
	return &CreateClassUseCase{ClassRepository: ClassRepository}
}

func (u *CreateClassUseCase) Execute(input CreateClassInputDto) (*CreateClassOutputDto, error) {

	// Create a new class entity
	newClass := entity.NewClass(input.TitleOfLesson, input.Description, input.TeacherID, input.StudentsIDs, input.GoogleMeetUrl)
	newClass.StudentsIDs = input.StudentsIDs

	// marshaled, _ := json.MarshalIndent(newClass, "", "\t")
	// fmt.Println(string(marshaled))

	// Save the new class using the repository
	if err := u.ClassRepository.Create(newClass); err != nil {
		return nil, err
	}

	// Return the output DTO with relevant information
	outputDto := &CreateClassOutputDto{
		ID:            newClass.ID,
		TitleOfLesson: newClass.TitleOfLesson,
		Description:   newClass.Description,
		TeacherID:     newClass.TeacherID,
		StudentsIDs:   newClass.StudentsIDs,
		GoogleMeetUrl: newClass.GoogleMeetUrl,
	}

	return outputDto, nil
}

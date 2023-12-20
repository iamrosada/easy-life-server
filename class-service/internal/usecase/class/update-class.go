package usecase

import "github.com/iamrosada/easy-life-server/class-service/internal/entity"

type UpdateClassInputDto struct {
	ID            string   `json:"id"`
	TitleOfLesson string   `json:"title_of_lesson"`
	Description   string   `json:"description"`
	TeacherID     string   `json:"teacher_id"`
	StudentsIDs   []string `json:"students_ids"`
	GoogleMeetUrl string   `json:"google_meet_url"`
}

type UpdateClassOutputDto struct {
	ID            string   `json:"id"`
	TitleOfLesson string   `json:"title_of_lesson"`
	Description   string   `json:"description"`
	TeacherID     string   `json:"teacher_id"`
	StudentsIDs   []string `json:"students_ids"`
	GoogleMeetUrl string   `json:"google_meet_url"`
}

type UpdateClassUseCase struct {
	ClassRepository entity.ClassRepository
}

func NewUpdateClassUseCase(ClassRepository entity.ClassRepository) *UpdateClassUseCase {
	return &UpdateClassUseCase{ClassRepository: ClassRepository}
}

func (u *UpdateClassUseCase) Execute(input UpdateClassInputDto) (*UpdateClassOutputDto, error) {
	// Get the existing class by ID
	existingClass, err := u.ClassRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	// Update the existing class with values from the input DTO
	existingClass.Update(input.TitleOfLesson, input.Description, input.TeacherID, input.GoogleMeetUrl)
	existingClass.StudentsIDs = input.StudentsIDs

	// Save the updated class using the repository
	err = u.ClassRepository.Update(existingClass)
	if err != nil {
		return nil, err
	}

	// Return the output DTO with relevant information
	outputDto := &UpdateClassOutputDto{
		ID:            existingClass.ID,
		TitleOfLesson: existingClass.TitleOfLesson,
		Description:   existingClass.Description,
		TeacherID:     existingClass.TeacherID,
		StudentsIDs:   existingClass.StudentsIDs,
		GoogleMeetUrl: existingClass.GoogleMeetUrl,
	}

	return outputDto, nil
}

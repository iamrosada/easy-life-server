package usecase

import "github.com/iamrosada/easy-life-server/class-service/internal/entity"

type GetClassByIDInputputDto struct {
	ID string `json:"id"`
}

type GetClassByIDOutputDto struct {
	ID            string   `json:"id"`
	TitleOfLesson string   `json:"title_of_lesson"`
	Description   string   `json:"description"`
	TeacherID     string   `json:"teacher_id"`
	GoogleMeetUrl string   `json:"google_meet_url"`
	StudentsIDs   []string `json:"students_ids"`
}

type GetClassByIDUseCase struct {
	ClassRepository entity.ClassRepository
}

func NewGetClassByIDUseCase(ClassRepository entity.ClassRepository) *GetClassByIDUseCase {
	return &GetClassByIDUseCase{ClassRepository: ClassRepository}
}

func (u *GetClassByIDUseCase) Execute(input GetClassByIDInputputDto) (*GetClassByIDOutputDto, error) {
	class, err := u.ClassRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	// Return the output DTO with relevant information
	outputDto := &GetClassByIDOutputDto{
		ID:            class.ID,
		TitleOfLesson: class.TitleOfLesson,
		Description:   class.Description,
		TeacherID:     class.TeacherID,
		StudentsIDs:   class.StudentsIDs,
		GoogleMeetUrl: class.GoogleMeetUrl,
	}

	return outputDto, nil
}

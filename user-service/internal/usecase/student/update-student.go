package student

import "github.com/iamrosada/easy-life-server/user-server/internal/entity"

type UpdateStudentInputDto struct {
	ID string `json:"id"`

	Name       string `json:"name"`
	FulName    string `json:"full_name"`
	CourseName string `json:"course_name"`
	EventID    string `json:"event_id"`
}

type UpdateStudentOutputDto struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	FulName    string `json:"full_name"`
	CourseName string `json:"course_name"`
	EventID    string `json:"event_id"`
}
type UpdateStudentUseCase struct {
	StudentRepository entity.StudentRepository
}

func NewUpdateStudentUseCase(StudentRepository entity.StudentRepository) *UpdateStudentUseCase {
	return &UpdateStudentUseCase{StudentRepository: StudentRepository}
}

func (u *UpdateStudentUseCase) Execute(input UpdateStudentInputDto) (*UpdateStudentOutputDto, error) {

	ExistStudent, err := u.StudentRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}
	ExistStudent.Update(input.CourseName, input.FulName, input.Name, input.EventID)

	err = u.StudentRepository.Update(ExistStudent)
	if err != nil {
		return nil, err
	}
	return &UpdateStudentOutputDto{
		ID:         ExistStudent.ID,
		FulName:    ExistStudent.FullName,
		Name:       ExistStudent.Name,
		CourseName: ExistStudent.CourseName,
		EventID:    ExistStudent.EventID,
	}, nil

}

package student

import "github.com/iamrosada/easy-life-server/user-server/internal/entity"

type GetStudentByEventIDInputputDto struct {
	EventID string `json:"event_id"`
}

type GetStudentByEventIDOutputDto struct {
	Students []*entity.Student `json:"students"`
}
type GetStudentByEventIDUseCase struct {
	StudentRepository entity.StudentRepository
}

func NewGetStudentByEventIDUseCase(StudentRepository entity.StudentRepository) *GetStudentByEventIDUseCase {
	return &GetStudentByEventIDUseCase{StudentRepository: StudentRepository}
}

func (u *GetStudentByEventIDUseCase) Execute(input GetStudentByEventIDInputputDto) (*GetStudentByEventIDOutputDto, error) {
	Student, err := u.StudentRepository.GetByEventID(input.EventID)
	if err != nil {
		return nil, err
	}
	return &GetStudentByEventIDOutputDto{Students: Student}, nil

}

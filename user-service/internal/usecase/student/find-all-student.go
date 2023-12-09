package student

import "github.com/iamrosada/easy-life-server/user-server/internal/entity"

type GetAllStudentOutputDto struct {
	Students []*entity.Student `json:"students"`
}
type GetAllStudentUseCase struct {
	StudentRepository entity.StudentRepository
}

func NewGetAllStudentUseCase(StudentRepository entity.StudentRepository) *GetAllStudentUseCase {
	return &GetAllStudentUseCase{StudentRepository: StudentRepository}
}

func (u *GetAllStudentUseCase) Execute() (*GetAllStudentOutputDto, error) {
	Students, err := u.StudentRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return &GetAllStudentOutputDto{Students: Students}, nil

}

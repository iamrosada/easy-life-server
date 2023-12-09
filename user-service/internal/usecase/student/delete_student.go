package student

import "github.com/iamrosada/easy-life-server/user-server/internal/entity"

type DeleteStudentInputDto struct {
	ID string `json:"id"`
}

type DeleteStudentOutputDto struct {
	ID string `json:"id"`
}
type DeleteStudentUseCase struct {
	StudentRepository entity.StudentRepository
}

func NewDeleteStudentUseCase(StudentRepository entity.StudentRepository) *DeleteStudentUseCase {
	return &DeleteStudentUseCase{StudentRepository: StudentRepository}

}

func (u *DeleteStudentUseCase) Execute(input DeleteStudentInputDto) (*DeleteStudentOutputDto, error) {

	err := u.StudentRepository.DeleteByID(input.ID)

	if err != nil {
		return nil, err
	}

	return &DeleteStudentOutputDto{ID: input.ID}, err

}

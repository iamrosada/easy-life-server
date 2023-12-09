package teacher

import "github.com/iamrosada/easy-life-server/user-server/internal/entity"

type DeleteTeacherInputDto struct {
	ID string `json:"id"`
}

type DeleteTeacherOutputDto struct {
	ID string `json:"id"`
}
type DeleteTeacherUseCase struct {
	TeacherRepository entity.TeacherRepository
}

func NewDeleteTeacherUseCase(TeacherRepository entity.TeacherRepository) *DeleteTeacherUseCase {
	return &DeleteTeacherUseCase{TeacherRepository: TeacherRepository}

}

func (u *DeleteTeacherUseCase) Execute(input DeleteTeacherInputDto) (*DeleteTeacherOutputDto, error) {

	err := u.TeacherRepository.DeleteByID(input.ID)

	if err != nil {
		return nil, err
	}

	return &DeleteTeacherOutputDto{ID: input.ID}, err

}

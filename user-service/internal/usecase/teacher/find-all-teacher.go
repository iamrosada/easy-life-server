package teacher

import "github.com/iamrosada/easy-life-server/user-server/internal/entity"

type GetAllTeacherOutputDto struct {
	Teachers []*entity.Teacher `json:"teachers"`
}
type GetAllTeacherUseCase struct {
	TeacherRepository entity.TeacherRepository
}

func NewGetAllTeacherUseCase(TeacherRepository entity.TeacherRepository) *GetAllTeacherUseCase {
	return &GetAllTeacherUseCase{TeacherRepository: TeacherRepository}
}

func (u *GetAllTeacherUseCase) Execute() (*GetAllTeacherOutputDto, error) {
	Teachers, err := u.TeacherRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return &GetAllTeacherOutputDto{Teachers: Teachers}, nil

}

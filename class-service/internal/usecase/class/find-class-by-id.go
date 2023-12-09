package usecase

import "github.com/iamrosada/easy-life-server/class-service/internal/entity"

type GetClassByIDInputputDto struct {
	ID string `json:"id"`
}

type GetClassByIDOutputDto struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	FullName   string `json:"full_name"`
	CourseName string `json:"course_name"`
}
type GetClassByIDUseCase struct {
	ClassRepository entity.ClassRepository
}

func NewGetClassByIDUseCase(ClassRepository entity.ClassRepository) *GetClassByIDUseCase {
	return &GetClassByIDUseCase{ClassRepository: ClassRepository}
}

func (u *GetClassByIDUseCase) Execute(input GetClassByIDInputputDto) (*GetClassByIDOutputDto, error) {
	Class, err := u.ClassRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}
	return &GetClassByIDOutputDto{
		ID:         Class.ID,
		FullName:   Class.FullName,
		Name:       Class.Name,
		CourseName: Class.CourseName,
	}, nil

}

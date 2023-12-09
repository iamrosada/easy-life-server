package usecase

import "github.com/iamrosada/easy-life-server/class-service/internal/entity"

type CreateClassInputDto struct {
	Name       string `json:"name"`
	FullName   string `json:"full_name"`
	CourseName string `json:"course_name"`
}

type CreateClassOutputDto struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	FulName    string `json:"full_name"`
	CourseName string `json:"course_name"`
}
type CreateClassUseCase struct {
	ClassRepository entity.ClassRepository
}

func NewCreateClassUseCase(ClassRepository entity.ClassRepository) *CreateClassUseCase {
	return &CreateClassUseCase{ClassRepository: ClassRepository}

}

func (u *CreateClassUseCase) Execute(input CreateClassInputDto) (*CreateClassOutputDto, error) {

	Class := entity.NewClass(
		input.CourseName,
		input.FullName,
		input.Name,
	)

	err := u.ClassRepository.Create(Class)

	if err != nil {
		return nil, err
	}

	return &CreateClassOutputDto{
		ID:         Class.ID,
		FulName:    Class.FullName,
		Name:       Class.Name,
		CourseName: Class.CourseName,
	}, nil
}

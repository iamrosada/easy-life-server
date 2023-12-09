package usecase

import "github.com/iamrosada/easy-life-server/class-service/internal/entity"

type UpdateClassInputDto struct {
	ID string `json:"id"`

	Name       string `json:"name"`
	FulName    string `json:"full_name"`
	CourseName string `json:"course_name"`
}

type UpdateClassOutputDto struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	FulName    string `json:"full_name"`
	CourseName string `json:"course_name"`
}
type UpdateClassUseCase struct {
	ClassRepository entity.ClassRepository
}

func NewUpdateClassUseCase(ClassRepository entity.ClassRepository) *UpdateClassUseCase {
	return &UpdateClassUseCase{ClassRepository: ClassRepository}
}

func (u *UpdateClassUseCase) Execute(input UpdateClassInputDto) (*UpdateClassOutputDto, error) {

	ExistClass, err := u.ClassRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}
	ExistClass.Update(input.CourseName, input.FulName, input.Name)

	err = u.ClassRepository.Update(ExistClass)
	if err != nil {
		return nil, err
	}
	return &UpdateClassOutputDto{
		ID:         ExistClass.ID,
		FulName:    ExistClass.FullName,
		Name:       ExistClass.Name,
		CourseName: ExistClass.CourseName,
	}, nil

}

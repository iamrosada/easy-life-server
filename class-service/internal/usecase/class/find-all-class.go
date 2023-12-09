package usecase

import "github.com/iamrosada/easy-life-server/class-service/internal/entity"

type GetAllClassOutputDto struct {
	Classs []*entity.Class `json:"classs"`
}
type GetAllClassUseCase struct {
	ClassRepository entity.ClassRepository
}

func NewGetAllClassUseCase(ClassRepository entity.ClassRepository) *GetAllClassUseCase {
	return &GetAllClassUseCase{ClassRepository: ClassRepository}
}

func (u *GetAllClassUseCase) Execute() (*GetAllClassOutputDto, error) {
	Classs, err := u.ClassRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return &GetAllClassOutputDto{Classs: Classs}, nil

}

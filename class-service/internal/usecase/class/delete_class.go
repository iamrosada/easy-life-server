package usecase

import "github.com/iamrosada/easy-life-server/class-service/internal/entity"

type DeleteClassInputDto struct {
	ID string `json:"id"`
}

type DeleteClassOutputDto struct {
	ID string `json:"id"`
}
type DeleteClassUseCase struct {
	ClassRepository entity.ClassRepository
}

func NewDeleteClassUseCase(ClassRepository entity.ClassRepository) *DeleteClassUseCase {
	return &DeleteClassUseCase{ClassRepository: ClassRepository}

}

func (u *DeleteClassUseCase) Execute(input DeleteClassInputDto) (*DeleteClassOutputDto, error) {

	err := u.ClassRepository.DeleteByID(input.ID)

	if err != nil {
		return nil, err
	}

	return &DeleteClassOutputDto{ID: input.ID}, err

}

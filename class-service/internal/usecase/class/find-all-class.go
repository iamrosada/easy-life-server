package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/iamrosada/easy-life-server/class-service/internal/entity"
)

type GetAllClassOutputDto struct {
	Class []*entity.Class `json:"classes"`
}
type GetAllClassUseCase struct {
	ClassRepository entity.ClassRepository
}

func NewGetAllClassUseCase(ClassRepository entity.ClassRepository) *GetAllClassUseCase {
	return &GetAllClassUseCase{ClassRepository: ClassRepository}
}

func (u *GetAllClassUseCase) Execute() (*GetAllClassOutputDto, error) {
	class, err := u.ClassRepository.FindAll()
	if err != nil {
		return nil, err
	}
	marshaled, _ := json.MarshalIndent(class, "", "\t")
	fmt.Println(string(marshaled))
	return &GetAllClassOutputDto{Class: class}, nil

}

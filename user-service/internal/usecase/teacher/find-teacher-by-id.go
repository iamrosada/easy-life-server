package teacher

import "github.com/iamrosada/easy-life-server/user-server/internal/entity"

type GetTeacherByIDInputputDto struct {
	ID string `json:"id"`
}

type GetTeacherByIDOutputDto struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	IsActive       bool   `json:"is_active"`
	CourseLanguage string `json:"course_language"`
}
type GetTeacherByIDUseCase struct {
	TeacherRepository entity.TeacherRepository
}

func NewGetTeacherByIDUseCase(TeacherRepository entity.TeacherRepository) *GetTeacherByIDUseCase {
	return &GetTeacherByIDUseCase{TeacherRepository: TeacherRepository}
}

func (u *GetTeacherByIDUseCase) Execute(input GetTeacherByIDInputputDto) (*GetTeacherByIDOutputDto, error) {
	Teacher, err := u.TeacherRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}
	return &GetTeacherByIDOutputDto{
		ID:             Teacher.ID,
		IsActive:       Teacher.IsActive,
		Name:           Teacher.Name,
		CourseLanguage: Teacher.CourseLanguage,
	}, nil

}

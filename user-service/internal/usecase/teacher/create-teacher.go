package teacher

import "github.com/iamrosada/easy-life-server/user-server/internal/entity"

type CreateTeacherInputDto struct {
	Name           string `json:"name"`
	IsActive       bool   `json:"is_active"`
	CourseLanguage string `json:"course_language"`
}

type CreateTeacherOutputDto struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	IsActive       bool   `json:"is_active"`
	CourseLanguage string `json:"course_language"`
}
type CreateTeacherUseCase struct {
	TeacherRepository entity.TeacherRepository
}

func NewCreateTeacherUseCase(TeacherRepository entity.TeacherRepository) *CreateTeacherUseCase {
	return &CreateTeacherUseCase{TeacherRepository: TeacherRepository}

}

func (u *CreateTeacherUseCase) Execute(input CreateTeacherInputDto) (*CreateTeacherOutputDto, error) {

	Teacher := entity.NewTeacher(
		input.Name,
		input.CourseLanguage,
		input.IsActive,
	)

	err := u.TeacherRepository.Create(Teacher)

	if err != nil {
		return nil, err
	}

	return &CreateTeacherOutputDto{
		ID:             Teacher.ID,
		IsActive:       Teacher.IsActive,
		Name:           Teacher.Name,
		CourseLanguage: Teacher.CourseLanguage,
	}, nil
}

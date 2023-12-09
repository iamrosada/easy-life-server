package teacher

import "github.com/iamrosada/easy-life-server/user-server/internal/entity"

type UpdateTeacherInputDto struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	IsActive       bool   `json:"is_active"`
	CourseLanguage string `json:"course_name"`
}

type UpdateTeacherOutputDto struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	IsActive       bool   `json:"is_active"`
	CourseLanguage string `json:"course_name"`
}
type UpdateTeacherUseCase struct {
	TeacherRepository entity.TeacherRepository
}

func NewUpdateTeacherUseCase(TeacherRepository entity.TeacherRepository) *UpdateTeacherUseCase {
	return &UpdateTeacherUseCase{TeacherRepository: TeacherRepository}
}

func (u *UpdateTeacherUseCase) Execute(input UpdateTeacherInputDto) (*UpdateTeacherOutputDto, error) {

	ExistTeacher, err := u.TeacherRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}
	ExistTeacher.Update(input.Name, input.CourseLanguage, input.IsActive)

	err = u.TeacherRepository.Update(ExistTeacher)
	if err != nil {
		return nil, err
	}
	return &UpdateTeacherOutputDto{
		ID:             ExistTeacher.ID,
		IsActive:       ExistTeacher.IsActive,
		Name:           ExistTeacher.Name,
		CourseLanguage: ExistTeacher.CourseLanguage,
	}, nil

}

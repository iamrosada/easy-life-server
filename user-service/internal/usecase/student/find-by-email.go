package student

import "github.com/iamrosada/easy-life-server/user-server/internal/entity"

type GetStudentByEmailInputputDto struct {
	Email string `json:"email"`
}

type GetStudentByEmailOutputDto struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	FullName   string `json:"full_name"`
	CourseName string `json:"course_name"`
	Email      string `json:"email"`
	EventID    string `json:"event_id"`

	TeachersIDs []string `json:"teachers_ids" gorm:"type:varchar[]"`

	// TeachersIDs []string `gorm:"type:jsonb" json:"teachers_ids"`
}
type GetStudentByEmailUseCase struct {
	StudentRepository entity.StudentRepository
}

func NewGetStudentByEmailUseCase(StudentRepository entity.StudentRepository) *GetStudentByEmailUseCase {
	return &GetStudentByEmailUseCase{StudentRepository: StudentRepository}
}

func (u *GetStudentByEmailUseCase) Execute(input GetStudentByEmailInputputDto) (*GetStudentByEmailOutputDto, error) {
	Student, err := u.StudentRepository.GetByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	return &GetStudentByEmailOutputDto{
		ID:          Student.ID,
		FullName:    Student.FullName,
		Name:        Student.Name,
		EventID:     Student.EventID,
		Email:       Student.Email,
		CourseName:  Student.CourseName,
		TeachersIDs: Student.TeachersIDs,
	}, nil

}

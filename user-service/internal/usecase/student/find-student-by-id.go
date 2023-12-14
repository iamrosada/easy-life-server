package student

import "github.com/iamrosada/easy-life-server/user-server/internal/entity"

type GetStudentByIDInputputDto struct {
	ID string `json:"id"`
}

type GetStudentByIDOutputDto struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	FullName   string `json:"full_name"`
	Email      string `json:"email"`
	CourseName string `json:"course_name"`
	EventID    string `json:"event_id"`

	TeachersIDs []string `json:"teachers_ids" gorm:"type:varchar[]"`

	// TeachersIDs []string `gorm:"type:jsonb" json:"teachers_ids"`
}
type GetStudentByIDUseCase struct {
	StudentRepository entity.StudentRepository
}

func NewGetStudentByIDUseCase(StudentRepository entity.StudentRepository) *GetStudentByIDUseCase {
	return &GetStudentByIDUseCase{StudentRepository: StudentRepository}
}

func (u *GetStudentByIDUseCase) Execute(input GetStudentByIDInputputDto) (*GetStudentByIDOutputDto, error) {
	Student, err := u.StudentRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}
	return &GetStudentByIDOutputDto{
		ID:          Student.ID,
		FullName:    Student.FullName,
		Name:        Student.Name,
		CourseName:  Student.CourseName,
		EventID:     Student.EventID,
		Email:       Student.Email,
		TeachersIDs: Student.TeachersIDs,
	}, nil

}

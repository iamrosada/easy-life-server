package student

import "github.com/iamrosada/easy-life-server/user-server/internal/entity"

type CreateStudentInputDto struct {
	Name        string   `json:"name"`
	FullName    string   `json:"full_name"`
	CourseName  string   `json:"course_name"`
	TeachersIDs []string `json:"teachers_ids" gorm:"type:varchar[]"`

	// TeachersIDs []string `gorm:"type:jsonb" json:"teachers_ids"`
}

type CreateStudentOutputDto struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	FullName    string   `json:"full_name"` // Corrected field name
	CourseName  string   `json:"course_name"`
	TeachersIDs []string `json:"teachers_ids" gorm:"type:varchar[]"`

	// TeachersIDs []string `gorm:"type:jsonb" json:"teachers_ids"`
}
type CreateStudentUseCase struct {
	StudentRepository entity.StudentRepository
}

func NewCreateStudentUseCase(StudentRepository entity.StudentRepository) *CreateStudentUseCase {
	return &CreateStudentUseCase{StudentRepository: StudentRepository}

}

func (u *CreateStudentUseCase) Execute(input CreateStudentInputDto) (*CreateStudentOutputDto, error) {
	// Create a new student
	student := entity.NewStudent(
		input.Name,
		input.FullName,
		input.CourseName,
	)

	// Set the TeachersIDs from the input
	student.TeachersIDs = input.TeachersIDs

	// Create the student in the repository
	err := u.StudentRepository.Create(student)
	if err != nil {
		return nil, err
	}

	// Return the output DTO
	return &CreateStudentOutputDto{
		ID:          student.ID,
		FullName:    student.FullName,
		Name:        student.Name,
		CourseName:  student.CourseName,
		TeachersIDs: student.TeachersIDs,
	}, nil
}

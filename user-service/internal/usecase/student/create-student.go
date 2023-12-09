package student

import "github.com/iamrosada/easy-life-server/user-server/internal/entity"

type CreateStudentInputDto struct {
	Name        string   `json:"name"`
	FullName    string   `json:"full_name"`
	CourseName  string   `json:"course_name"`
	TeachersIDs []string `json:"teachers_ids"`
}

type CreateStudentOutputDto struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	FullName    string   `json:"full_name"` // Corrected field name
	CourseName  string   `json:"course_name"`
	TeachersIDs []string `json:"teachers_ids"`
}
type CreateStudentUseCase struct {
	StudentRepository entity.StudentRepository
}

func NewCreateStudentUseCase(StudentRepository entity.StudentRepository) *CreateStudentUseCase {
	return &CreateStudentUseCase{StudentRepository: StudentRepository}

}

func (u *CreateStudentUseCase) Execute(input CreateStudentInputDto) (*CreateStudentOutputDto, error) {

	Student := entity.NewStudent(
		input.Name,
		input.FullName,
		input.CourseName,
	)

	err := u.StudentRepository.Create(Student)

	if err != nil {
		return nil, err
	}

	return &CreateStudentOutputDto{
		ID:          Student.ID,
		FullName:    Student.FullName,
		Name:        Student.Name,
		CourseName:  Student.CourseName,
		TeachersIDs: input.TeachersIDs,
	}, nil
}

package student

import "github.com/iamrosada/easy-life-server/user-server/internal/entity"

type CreateStudentInputDto struct {
	Name       string `json:"name"`
	FullName   string `json:"full_name"`
	CourseName string `json:"course_name"`
}

type CreateStudentOutputDto struct {
	ID         string         `json:"id"`
	Name       string         `json:"name"`
	FulName    string         `json:"full_name"`
	CourseName string         `json:"course_name"`
	Teacher    entity.Teacher `json:"teacher_ids"`
}
type CreateStudentUseCase struct {
	StudentRepository entity.StudentRepository
}

func NewCreateStudentUseCase(StudentRepository entity.StudentRepository) *CreateStudentUseCase {
	return &CreateStudentUseCase{StudentRepository: StudentRepository}

}

func (u *CreateStudentUseCase) Execute(input CreateStudentInputDto) (*CreateStudentOutputDto, error) {

	Student := entity.NewStudent(
		input.CourseName,
		input.FullName,
		input.Name,
	)

	err := u.StudentRepository.Create(Student)

	if err != nil {
		return nil, err
	}

	return &CreateStudentOutputDto{
		ID:         Student.ID,
		FulName:    Student.FullName,
		Name:       Student.Name,
		CourseName: Student.CourseName,
	}, nil
}

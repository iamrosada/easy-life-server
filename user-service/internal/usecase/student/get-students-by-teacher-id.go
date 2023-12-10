package student

import "github.com/iamrosada/easy-life-server/user-server/internal/entity"

type ListStudentsByTeacherIDUseCase struct {
	StudentRepository entity.StudentRepository
}

func NewListStudentsByTeacherIDUseCase(StudentRepository entity.StudentRepository) *ListStudentsByTeacherIDUseCase {
	return &ListStudentsByTeacherIDUseCase{StudentRepository: StudentRepository}

}

func (uc *ListStudentsByTeacherIDUseCase) ListStudentsByTeacherID(teacherId string) ([]*entity.Student, error) {

	students, err := uc.StudentRepository.ListStudentsByTeacherID(teacherId)
	if err != nil {
		return nil, err
	}

	return students, nil
}

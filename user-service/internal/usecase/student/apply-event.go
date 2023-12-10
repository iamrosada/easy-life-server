package student

import "github.com/iamrosada/easy-life-server/user-server/internal/entity"

type ApplyEventStudentUseCase struct {
	StudentRepository entity.StudentRepository
}

func NewCreateEventStudentUseCase(StudentRepository entity.StudentRepository) *ApplyEventStudentUseCase {
	return &ApplyEventStudentUseCase{StudentRepository: StudentRepository}

}

func (uc *ApplyEventStudentUseCase) ApplyEvent(eventID string, userIDs []string) error {
	return uc.StudentRepository.ApplyEvent(eventID, userIDs)

}

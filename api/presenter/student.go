package presenter

import "github.com/omidnasiri/mediana-sms/internal/models"

type CreateStudentRequestDTO struct {
	Name     string `json:"name" binding:"required"`
	SchoolId uint   `json:"school_id" binding:"required"`
}

type BulkAddStudentToTeacherRequestDTO struct {
	StudentIds []uint `json:"student_ids" binding:"required"`
}

type BulkAddStudentToTeacherResponseDTO struct {
	AffectedStudentsCount int               `json:"affected_students_count"`
	Students              []*models.Student `json:"students"`
}

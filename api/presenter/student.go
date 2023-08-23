package presenter

type CreateStudentRequestDTO struct {
	Name     string `json:"name" binding:"required"`
	SchoolId uint   `json:"school_id" binding:"required"`
}

package presenter

type CreateTeacherRequestDTO struct {
	Name     string `json:"name" binding:"required"`
	SchoolId uint   `json:"school_id" binding:"required"`
}

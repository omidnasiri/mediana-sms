package presenter

type CreateTeacherRequestDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	SchoolId uint   `json:"school_id" binding:"required"`
}

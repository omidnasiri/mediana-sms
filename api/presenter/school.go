package presenter

type CreateSchoolRequestDTO struct {
	Title        string `json:"title" binding:"required"`
	HeadmasterId uint   `json:"headmaster_id" binding:"required"`
}

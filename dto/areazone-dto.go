package dto

type AreazoneCreateDto struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description"`
	Created_by  uint64 `json:"created_by"`
	Status      uint64 `json:"status" form:"status"`
}
type AreazoneUpdateDto struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description"`
	Created_by  uint64 `json:"created_by"`
	Status      uint64 `json:"status" form:"status"`
}

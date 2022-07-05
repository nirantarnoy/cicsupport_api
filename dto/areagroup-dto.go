package dto

type AreagroupUpdateDTO struct {
	ID         uint64 `json:"id" form:"id" binding:"required"`
	Name       string `json:"name" form:"name" binding:"required"`
	Created_by uint64 `json:"created_by,omitempty" form:"created_by,omitempty"`
	Status     uint64 `json:"status" form:"status"`
}
type AreagroupDTO struct {
	Name       string `json:"name" form:"name" binding:"required"`
	Created_by uint64 `json:"created_by,omitempty" form:"created_by,omitempty"`
	Status     uint64 `json:"status" form:"status"`
}

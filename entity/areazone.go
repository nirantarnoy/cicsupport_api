package entity

type Areazone struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Created_by  uint64 `json:"created_by"`
	Status      uint64 `json:"status"`
}

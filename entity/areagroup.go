package entity

type Areagroup struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	Created_by uint64 `json:"created_by"`
	Status     uint64 `json:"status"`
}

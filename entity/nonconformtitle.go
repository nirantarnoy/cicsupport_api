package entity

type NonConformTitle struct {
	Id           string `json:"id"`
	Code         string `json:"code"`
	Name         string `json:"name"`
	ModuleTypeId string `json:"module_type_id"`
	Status       string `json:"status"`
}

package entity

type TeamNotify struct {
	Id           string `json:"id"`
	TransRefId   string `json:"trans_ref_id"`
	ModuleTypeId string `json:"module_type_id"`
	EmpId        string `json:"emp_id"`
	Title        string `json:"title"`
	Detail       string `json:"detail"`
	ReadStatus   string `json:"read_status"`
	NotifyDate   string `json:"notify_date"`
}

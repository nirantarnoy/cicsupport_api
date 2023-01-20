package entity

type InspectionTransHistory struct {
	PlanId         string `json:"plan_id"`
	PlanNo         string `json:"plan_no"`
	TeamId         string `json:"team_id"`
	PlanDate       string `json:"plan_date"`
	PlanActualDate string `json:"plan_actual_date"`
	Status         uint64 `json:"status"`
}

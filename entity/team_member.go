package entity

type TeamMember struct {
	CurrentTeamId string `json:"currnet_team_id"`
	Fname         string `json:"fname"`
	Lname         string `json:"lname"`
	IsHead        string `json:"is_head"`
}

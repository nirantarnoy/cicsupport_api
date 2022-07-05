package entity

type User struct {
	ID     uint64 `json:"id"`
	AdUser string `json:"ad_user"`
	Token  string `gorm:"-" json:"token,omitempty"`
}

package entity

type User struct {
	ID       uint64 `json:"id"`
	Dns_user string `json:"dns_user"`
	Token    string `gorm:"-" json:"token,omitempty"`
}

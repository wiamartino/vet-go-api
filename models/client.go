package models

type Client struct {
	ClientID  uint   `gorm:"primaryKey" json:"client_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

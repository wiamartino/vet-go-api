package domain

type Client struct {
	ClientID  uint   `gorm:"primaryKey" json:"client_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Pets      []Pet  `gorm:"foreignKey:ClientID" json:"pets"`
}

type ClientRepository interface {
	FindAll() ([]Client, error)
	FindByID(id uint) (Client, error)
	Create(client *Client) error
	Update(client *Client) error
	Delete(id uint) error
}

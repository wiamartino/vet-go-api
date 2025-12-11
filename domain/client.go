package domain

type Client struct {
	ClientID  uint   `gorm:"primaryKey" json:"client_id"`
	FirstName string `json:"first_name" binding:"required,min=2,max=100"`
	LastName  string `json:"last_name" binding:"required,min=2,max=100"`
	Address   string `json:"address" binding:"required,min=5,max=255"`
	Phone     string `json:"phone" binding:"required,min=10,max=20"`
	Email     string `json:"email" binding:"required,email"`
	Pets      []Pet  `gorm:"foreignKey:ClientID" json:"pets"`
}

type ClientRepository interface {
	FindAll() ([]Client, error)
	FindByID(id uint) (Client, error)
	Create(client *Client) error
	Update(client *Client) error
	Delete(id uint) error
}

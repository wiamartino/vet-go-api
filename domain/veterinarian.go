package domain

type Veterinarian struct {
	VeterinarianID uint   `gorm:"primaryKey" json:"veterinarian_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Specialty      string `json:"specialty"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
}

type VeterinarianRepository interface {
	FindAll() ([]Veterinarian, error)
	FindByID(id uint) (Veterinarian, error)
	Create(veterinarian *Veterinarian) error
	Update(veterinarian *Veterinarian) error
	Delete(id uint) error
}

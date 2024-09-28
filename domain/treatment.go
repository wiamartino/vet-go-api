package domain

type Treatment struct {
	TreatmentID uint    `gorm:"primaryKey" json:"treatment_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Cost        float64 `json:"cost"`
}

type TreatmentRepository interface {
	FindAll() ([]Treatment, error)
	FindByID(id uint) (Treatment, error)
	Create(treatment *Treatment) error
	Update(treatment *Treatment) error
	Delete(id uint) error
}

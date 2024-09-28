package domain

type Medication struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type MedicationRepository interface {
	FindAll() ([]Medication, error)
	FindByID(id uint) (Medication, error)
	Create(medication Medication) error
	Update(medication Medication) error
	Delete(id uint) error
}

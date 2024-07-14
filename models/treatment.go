package models

type Treatment struct {
	TreatmentID uint    `gorm:"primaryKey" json:"treatment_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Cost        float64 `json:"cost"`
}

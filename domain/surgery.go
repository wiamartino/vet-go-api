package domain

import "time"

type SurgeryType string

const (
	SurgeryTypeRoutine   SurgeryType = "routine"
	SurgeryTypeEmergency SurgeryType = "emergency"
	SurgeryTypeElective  SurgeryType = "elective"
)

type SurgeryStatus string

const (
	SurgeryStatusScheduled  SurgeryStatus = "scheduled"
	SurgeryStatusInProgress SurgeryStatus = "in_progress"
	SurgeryStatusCompleted  SurgeryStatus = "completed"
	SurgeryStatusCancelled  SurgeryStatus = "cancelled"
)

type Surgery struct {
	SurgeryID        uint          `gorm:"primaryKey" json:"surgery_id"`
	PetID            uint          `json:"pet_id"`
	VeterinarianID   uint          `json:"veterinarian_id"`
	SurgeryName      string        `json:"surgery_name"`
	SurgeryType      SurgeryType   `json:"surgery_type" gorm:"default:'routine'"`
	Status           SurgeryStatus `json:"status" gorm:"default:'scheduled'"`
	ScheduledDate    time.Time     `json:"scheduled_date"`
	ActualDate       *time.Time    `json:"actual_date,omitempty"`
	Duration         *int          `json:"duration,omitempty"`
	PreOpNotes       string        `json:"pre_op_notes,omitempty"`
	PostOpNotes      string        `json:"post_op_notes,omitempty"`
	Complications    string        `json:"complications,omitempty"`
	AnesthesiaUsed   string        `json:"anesthesia_used,omitempty"`
	FollowUpRequired bool          `json:"follow_up_required" gorm:"default:true"`
	FollowUpDate     *time.Time    `json:"follow_up_date,omitempty"`
	Cost             *float64      `json:"cost,omitempty"`
	CreatedAt        time.Time     `json:"created_at"`
	UpdatedAt        time.Time     `json:"updated_at"`
	Pet              Pet           `gorm:"foreignKey:PetID" json:"pet,omitempty"`
	Veterinarian     Veterinarian  `gorm:"foreignKey:VeterinarianID" json:"veterinarian,omitempty"`
}

type SurgeryRepository interface {
	FindAll() ([]Surgery, error)
	FindByID(id uint) (Surgery, error)
	FindByPetID(petID uint) ([]Surgery, error)
	FindByStatus(status SurgeryStatus) ([]Surgery, error)
	FindScheduledSurgeries(startDate, endDate time.Time) ([]Surgery, error)
	Create(surgery *Surgery) error
	Update(surgery *Surgery) error
	Delete(id uint) error
}

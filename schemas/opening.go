package schemas

import "gorm.io/gorm"

type WorkType string

const (
	Onsite WorkType = "Onsite"
	Hybrid WorkType = "Hybrid"
	Remote WorkType = "Remote"
)

// Opening represents a job opening with various attributes
type Opening struct {
	gorm.Model
	Role     string   `gorm:"not null"`                                                 // Role of the job
	Company  string   `gorm:"not null"`                                                 // Company offering the job
	Location string   `gorm:"not null"`                                                 // Location of the job
	WorkType WorkType `gorm:"type:enum('Onsite', 'Hybrid', 'Remote');default:'Onsite'"` // Type of work
	Link     string   `gorm:"not null"`                                                 // Link to job posting
	Salary   int64    `gorm:"default:0"`                                                // Salary offered
}

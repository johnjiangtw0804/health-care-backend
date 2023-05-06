package model

import (
	"time"
)

type PatientDashboardView struct {
	ID                   int
	FirstName            string `gorm:""`
	LastName             string
	Age                  int
	Sex                  string
	BloodType            string
	DOB                  time.Time
	AssignedDoctorID     int
	BodyTempertature     float64
	PulseRate            int
	RespirationRate      int
	SystolicPressure     int
	DiastolicPressure    int
	CurrentPrescribedMed string
	CurrentDisease       string
}

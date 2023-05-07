package model

import (
	"time"
)

type PatientDashboardView struct {
	ID                      int
	FirstName               string
	LastName                string
	Age                     int
	Sex                     string
	BloodType               string
	DOB                     time.Time
	AssignedDoctorID        int
	AssignedDoctorFirstName string
	AssignedDoctorLastName  string
	BodyTemperature         float64
	PulseRate               int
	RespirationRate         int
	SystolicPressure        int
	DiastolicPressure       int
	CurrentPrescribedMed    string
	CurrentDisease          string
}

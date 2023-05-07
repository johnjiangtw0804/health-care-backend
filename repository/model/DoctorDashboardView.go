package model

import (
	"time"
)

type DoctorDashboardView struct {
	PatientID               int
	FirstName               string
	LastName                string
	Age                     int
	Sex                     string
	BloodType               string
	PhoneNumber             string
	Address                 string
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

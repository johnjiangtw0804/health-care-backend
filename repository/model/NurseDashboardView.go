package model

import (
	"time"
)

type NurseDashboardView struct {
	NurseID              int
	NurseFirstName       string
	NurseLastName        string
	PatientID            int
	PatientFirstName     string
	PatientLastName      string
	Age                  int
	Sex                  string
	BloodType            string
	PhoneNumber          string
	Address              string
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

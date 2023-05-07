package routes

import (
	repository "health-care-backend/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DashboardHandler struct {
	logger *zap.Logger
	repo   repository.Dashboard
}

func NewDashboardHandler(logger *zap.Logger, repo repository.Dashboard) *DashboardHandler {
	return &DashboardHandler{
		logger: logger,
		repo:   repo,
	}
}

type PatientDashboardResp struct {
	ID                      int          `json:"patient_id"`
	FirstName               string       `json:"first_name"`
	LastName                string       `json:"last_name"`
	Age                     int          `json:"age"`
	Sex                     string       `json:"sex"`
	BloodType               string       `json:"blood_type"`
	DOB                     time.Time    `json:"dob"`
	AssignedDoctorID        int          `json:"assigned_doctor_id"`
	AssignedDoctorFirstName string       `json:"assigned_doctor_first_name"`
	AssignedDoctorLastName  string       `json:"assigned_doctor_last_name"`
	BodyTemperature         float64      `json:"body_temperature"`
	PulseRate               int          `json:"pulse_rate"`
	RespirationRate         int          `json:"respiration_rate"`
	SystolicPressure        int          `json:"systolic_pressure"`
	DiastolicPressure       int          `json:"diastolic_pressure"`
	CurrentPrescribedMeds   []Medication `json:"current_prescribed_meds"`
	CurrentDiseases         []Disease    `json:"current_diseases"`
}

func (h *DashboardHandler) GetPatientDashboard(ctx *gin.Context) {
	pidStr := ctx.Query("patient_id")
	if pidStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "patient_id is required"})
		return
	}
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "patient_id must be an integer"})
		return
	}
	patientViews, err := h.repo.SelectPatientDashboard(pid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	type tempPatient struct {
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
		CurrentPrescribedMeds   map[string]int
		CurrentDiseases         map[string]int
	}
	idToPatient := make(map[int]tempPatient)
	for _, view := range patientViews {
		if _, ok := idToPatient[view.ID]; !ok {
			idToPatient[view.ID] = tempPatient{
				ID:                      view.ID,
				FirstName:               view.FirstName,
				LastName:                view.LastName,
				Age:                     view.Age,
				Sex:                     view.Sex,
				BloodType:               view.BloodType,
				DOB:                     view.DOB,
				AssignedDoctorID:        view.AssignedDoctorID,
				AssignedDoctorFirstName: view.AssignedDoctorFirstName,
				AssignedDoctorLastName:  view.AssignedDoctorLastName,
				BodyTemperature:         view.BodyTemperature,
				PulseRate:               view.PulseRate,
				RespirationRate:         view.RespirationRate,
				SystolicPressure:        view.SystolicPressure,
				DiastolicPressure:       view.DiastolicPressure,
				CurrentPrescribedMeds:   make(map[string]int),
				CurrentDiseases:         make(map[string]int),
			}
		}
		p := idToPatient[view.ID]
		p.CurrentPrescribedMeds[view.CurrentPrescribedMed] = 1
		p.CurrentDiseases[view.CurrentDisease] = 1
		idToPatient[view.ID] = p
	}
	var resp PatientDashboardResp
	for _, patient := range idToPatient {
		// we can only have one patient since the pid is unique
		resp = PatientDashboardResp{
			ID:                      patient.ID,
			FirstName:               patient.FirstName,
			LastName:                patient.LastName,
			Age:                     patient.Age,
			Sex:                     patient.Sex,
			BloodType:               patient.BloodType,
			DOB:                     patient.DOB,
			AssignedDoctorID:        patient.AssignedDoctorID,
			AssignedDoctorFirstName: patient.AssignedDoctorFirstName,
			AssignedDoctorLastName:  patient.AssignedDoctorLastName,
			BodyTemperature:         patient.BodyTemperature,
			PulseRate:               patient.PulseRate,
			RespirationRate:         patient.RespirationRate,
			SystolicPressure:        patient.SystolicPressure,
			DiastolicPressure:       patient.DiastolicPressure,
		}
		// convert map to array
		for med := range patient.CurrentPrescribedMeds {
			resp.CurrentPrescribedMeds = append(resp.CurrentPrescribedMeds, Medication{
				Name: med,
			})
		}
		for disease := range patient.CurrentDiseases {
			resp.CurrentDiseases = append(resp.CurrentDiseases, Disease{
				Name: disease,
			})
		}
	}
	ctx.JSON(http.StatusOK, resp)
}

type NurseDashboardResp struct {
	Patients []NursePatient `json:"patients"`
}
type NursePatient struct {
	NurseID                 int          `json:"nurse_id"`
	NurseFirstName          string       `json:"nurse_first_name"`
	NurseLastName           string       `json:"nurse_last_name"`
	PatientID               int          `json:"patient_id"`
	PatientFirstName        string       `json:"patient_first_name"`
	PatientLastName         string       `json:"patient_last_name"`
	Age                     int          `json:"age"`
	Sex                     string       `json:"sex"`
	BloodType               string       `json:"blood_type"`
	PhoneNumber             string       `json:"phone_number"`
	Address                 string       `json:"address"`
	DOB                     time.Time    `json:"dob"`
	AssignedDoctorID        int          `json:"assigned_doctor_id"`
	AssignedDoctorFirstName string       `json:"assigned_doctor_first_name"`
	AssignedDoctorLastName  string       `json:"assigned_doctor_last_name"`
	BodyTemperature         float64      `json:"body_temperature"`
	PulseRate               int          `json:"pulse_rate"`
	RespirationRate         int          `json:"respiration_rate"`
	SystolicPressure        int          `json:"systolic_pressure"`
	DiastolicPressure       int          `json:"diastolic_pressure"`
	CurrentPrescribedMeds   []Medication `json:"current_prescribed_meds"`
	CurrentDiseases         []Disease    `json:"current_diseases"`
}

func (h *DashboardHandler) GetNurseDashboard(ctx *gin.Context) {
	nidStr := ctx.Query("nurse_id")
	if nidStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "nurse_id is required"})
		return
	}
	nid, err := strconv.Atoi(nidStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "nurse_id must be integer"})
		return
	}
	views, err := h.repo.SelectNurseDashboard(nid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	type tempNursePatient struct {
		NurseID                 int
		NurseFirstName          string
		NurseLastName           string
		PatientID               int
		PatientFirstName        string
		PatientLastName         string
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
		CurrentPrescribedMeds   map[string]int
		CurrentDiseases         map[string]int
	}

	idToPatient := make(map[int]tempNursePatient)
	for _, view := range views {
		if _, ok := idToPatient[view.PatientID]; !ok {
			idToPatient[view.PatientID] = tempNursePatient{
				NurseID:                 view.NurseID,
				NurseFirstName:          view.NurseFirstName,
				NurseLastName:           view.NurseLastName,
				PatientID:               view.PatientID,
				PatientFirstName:        view.PatientFirstName,
				PatientLastName:         view.PatientLastName,
				Age:                     view.Age,
				Sex:                     view.Sex,
				BloodType:               view.BloodType,
				PhoneNumber:             view.PhoneNumber,
				Address:                 view.Address,
				DOB:                     view.DOB,
				AssignedDoctorID:        view.AssignedDoctorID,
				AssignedDoctorFirstName: view.AssignedDoctorFirstName,
				AssignedDoctorLastName:  view.AssignedDoctorLastName,
				BodyTemperature:         view.BodyTemperature,
				PulseRate:               view.PulseRate,
				RespirationRate:         view.RespirationRate,
				SystolicPressure:        view.SystolicPressure,
				DiastolicPressure:       view.DiastolicPressure,
				CurrentPrescribedMeds:   make(map[string]int),
				CurrentDiseases:         make(map[string]int),
			}
		}
		p := idToPatient[view.PatientID]
		p.CurrentPrescribedMeds[view.CurrentPrescribedMed] = 1
		p.CurrentDiseases[view.CurrentDisease] = 1
		idToPatient[view.PatientID] = p
	}

	var resp NurseDashboardResp
	for _, patient := range idToPatient {
		patientResp := NursePatient{
			NurseID:                 patient.NurseID,
			NurseFirstName:          patient.NurseFirstName,
			NurseLastName:           patient.NurseLastName,
			PatientID:               patient.PatientID,
			PatientFirstName:        patient.PatientFirstName,
			PatientLastName:         patient.PatientLastName,
			Age:                     patient.Age,
			Sex:                     patient.Sex,
			BloodType:               patient.BloodType,
			PhoneNumber:             patient.PhoneNumber,
			Address:                 patient.Address,
			DOB:                     patient.DOB,
			AssignedDoctorID:        patient.AssignedDoctorID,
			AssignedDoctorFirstName: patient.AssignedDoctorFirstName,
			AssignedDoctorLastName:  patient.AssignedDoctorLastName,
			BodyTemperature:         patient.BodyTemperature,
			PulseRate:               patient.PulseRate,
			RespirationRate:         patient.RespirationRate,
			SystolicPressure:        patient.SystolicPressure,
			DiastolicPressure:       patient.DiastolicPressure,
		}
		for med := range patient.CurrentPrescribedMeds {
			patientResp.CurrentPrescribedMeds = append(patientResp.CurrentPrescribedMeds, Medication{
				Name: med,
			})
		}
		for disease := range patient.CurrentDiseases {
			patientResp.CurrentDiseases = append(patientResp.CurrentDiseases, Disease{
				Name: disease,
			})
		}
		resp.Patients = append(resp.Patients, patientResp)
	}

	ctx.JSON(http.StatusOK, resp)
}

type DoctorDashboardResp struct {
	Patients []DoctorPatient `json:"patients"`
}
type DoctorPatient struct {
	PatientID               int          `json:"patient_id"`
	FirstName               string       `json:"first_name"`
	LastName                string       `json:"last_name"`
	Age                     int          `json:"age"`
	Sex                     string       `json:"sex"`
	BloodType               string       `json:"blood_type"`
	PhoneNumber             string       `json:"phone_number"`
	Address                 string       `json:"address"`
	DOB                     time.Time    `json:"dob"`
	AssignedDoctorID        int          `json:"assigned_doctor_id"`
	AssignedDoctorFirstName string       `json:"assigned_doctor_first_name"`
	AssignedDoctorLastName  string       `json:"assigned_doctor_last_name"`
	BodyTemperature         float64      `json:"body_temperature"`
	PulseRate               int          `json:"pulse_rate"`
	RespirationRate         int          `json:"respiration_rate"`
	SystolicPressure        int          `json:"systolic_pressure"`
	DiastolicPressure       int          `json:"diastolic_pressure"`
	CurrentPrescribedMeds   []Medication `json:"current_prescribed_meds"`
	CurrentDiseases         []Disease    `json:"current_diseases"`
}
type Medication struct {
	Name string `json:"name"`
}

type Disease struct {
	Name string `json:"name"`
}

func (h *DashboardHandler) GetDoctorDashboard(ctx *gin.Context) {
	didStr := ctx.Query("doctor_id")
	if didStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "doctor_id is required"})
		return
	}
	did, err := strconv.Atoi(didStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "doctor_id must be integer"})
		return
	}
	views, err := h.repo.SelectDoctorDashboard(did)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	type tempDoctorPatient struct {
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
		CurrentPrescribedMeds   map[string]int
		CurrentDiseases         map[string]int
	}

	idToPatient := make(map[int]tempDoctorPatient)
	for _, view := range views {
		if _, ok := idToPatient[view.PatientID]; !ok {
			idToPatient[view.PatientID] = tempDoctorPatient{
				PatientID:               view.PatientID,
				FirstName:               view.FirstName,
				LastName:                view.LastName,
				Age:                     view.Age,
				Sex:                     view.Sex,
				BloodType:               view.BloodType,
				PhoneNumber:             view.PhoneNumber,
				Address:                 view.Address,
				DOB:                     view.DOB,
				AssignedDoctorID:        view.AssignedDoctorID,
				AssignedDoctorFirstName: view.AssignedDoctorFirstName,
				AssignedDoctorLastName:  view.AssignedDoctorLastName,
				BodyTemperature:         view.BodyTemperature,
				PulseRate:               view.PulseRate,
				RespirationRate:         view.RespirationRate,
				SystolicPressure:        view.SystolicPressure,
				DiastolicPressure:       view.DiastolicPressure,
				CurrentPrescribedMeds:   make(map[string]int),
				CurrentDiseases:         make(map[string]int),
			}
		}
		v := idToPatient[view.PatientID]
		v.CurrentPrescribedMeds[view.CurrentPrescribedMed] = 1
		v.CurrentDiseases[view.CurrentDisease] = 1
		idToPatient[view.PatientID] = v
	}

	var resp DoctorDashboardResp
	for _, patient := range idToPatient {
		patientResp := DoctorPatient{
			PatientID:               patient.PatientID,
			FirstName:               patient.FirstName,
			LastName:                patient.LastName,
			Age:                     patient.Age,
			Sex:                     patient.Sex,
			BloodType:               patient.BloodType,
			PhoneNumber:             patient.PhoneNumber,
			Address:                 patient.Address,
			DOB:                     patient.DOB,
			AssignedDoctorID:        patient.AssignedDoctorID,
			AssignedDoctorFirstName: patient.AssignedDoctorFirstName,
			AssignedDoctorLastName:  patient.AssignedDoctorLastName,
			BodyTemperature:         patient.BodyTemperature,
			PulseRate:               patient.PulseRate,
			RespirationRate:         patient.RespirationRate,
			SystolicPressure:        patient.SystolicPressure,
			DiastolicPressure:       patient.DiastolicPressure,
		}
		for med := range patient.CurrentPrescribedMeds {
			patientResp.CurrentPrescribedMeds = append(patientResp.CurrentPrescribedMeds, Medication{
				Name: med,
			})
		}
		for disease := range patient.CurrentDiseases {
			patientResp.CurrentDiseases = append(patientResp.CurrentDiseases, Disease{
				Name: disease,
			})
		}
		resp.Patients = append(resp.Patients, patientResp)
	}
	ctx.JSON(http.StatusOK, resp)
}

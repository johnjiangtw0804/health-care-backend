package routes

import (
	repository "health-care-backend/repository"
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
	ID                   int       `json:"patient_id"`
	FirstName            string    `json:"first_name"`
	LastName             string    `json:"last_name"`
	Age                  int       `json:"age"`
	Sex                  string    `json:"sex"`
	BloodType            string    `json:"blood_type"`
	DOB                  time.Time `json:"dob"`
	AssignedDoctorID     int       `json:"assigned_doctor_id"`
	BodyTemperature      float64   `json:"body_temperature"`
	PulseRate            int       `json:"pulse_rate"`
	RespirationRate      int       `json:"respiration_rate"`
	SystolicPressure     int       `json:"systolic_pressure"`
	DiastolicPressure    int       `json:"diastolic_pressure"`
	CurrentPrescribedMed string    `json:"current_prescribed_med"`
	CurrentDisease       string    `json:"current_disease"`
}

func (h *DashboardHandler) GetPatientDashboard(ctx *gin.Context) {
	pidStr := ctx.Query("patient_id")
	if pidStr == "" {
		ctx.JSON(400, gin.H{"error": "patient_id is required"})
		return
	}
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "patient_id must be an integer"})
		return
	}
	dashboard, err := h.repo.SelectPatientDashboard(pid)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	resp := PatientDashboardResp{
		ID:                   dashboard.ID,
		FirstName:            dashboard.FirstName,
		LastName:             dashboard.LastName,
		Age:                  dashboard.Age,
		Sex:                  dashboard.Sex,
		BloodType:            dashboard.BloodType,
		DOB:                  dashboard.DOB,
		AssignedDoctorID:     dashboard.AssignedDoctorID,
		BodyTemperature:      dashboard.BodyTempertature,
		PulseRate:            dashboard.PulseRate,
		RespirationRate:      dashboard.RespirationRate,
		SystolicPressure:     dashboard.SystolicPressure,
		DiastolicPressure:    dashboard.DiastolicPressure,
		CurrentPrescribedMed: dashboard.CurrentPrescribedMed,
		CurrentDisease:       dashboard.CurrentDisease,
	}
	ctx.JSON(200, resp)
}

type NurseDashboardResp struct {
	Patients []NursePatient `json:"patients"`
}
type NursePatient struct {
	NurseID              int       `json:"nurse_id"`
	NurseFirstName       string    `json:"nurse_first_name"`
	NurseLastName        string    `json:"nurse_last_name"`
	PatientID            int       `json:"patient_id"`
	PatientFirstName     string    `json:"patient_first_name"`
	PatientLastName      string    `json:"patient_last_name"`
	Age                  int       `json:"age"`
	Sex                  string    `json:"sex"`
	BloodType            string    `json:"blood_type"`
	PhoneNumber          string    `json:"phone_number"`
	Address              string    `json:"address"`
	DOB                  time.Time `json:"dob"`
	AssignedDoctorID     int       `json:"assigned_doctor_id"`
	BodyTempertature     float64   `json:"body_temperature"`
	PulseRate            int       `json:"pulse_rate"`
	RespirationRate      int       `json:"respiration_rate"`
	SystolicPressure     int       `json:"systolic_pressure"`
	DiastolicPressure    int       `json:"diastolic_pressure"`
	CurrentPrescribedMed string    `json:"current_prescribed_med"`
	CurrentDisease       string    `json:"current_disease"`
}

func (h *DashboardHandler) GetNurseDashboard(ctx *gin.Context) {
	nidStr := ctx.Query("nurse_id")
	if nidStr == "" {
		ctx.JSON(400, gin.H{"error": "nurse_id is required"})
		return
	}
	nid, err := strconv.Atoi(nidStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "nurse_id must be integer"})
		return
	}
	patients, err := h.repo.SelectNurseDashboard(nid)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var resp NurseDashboardResp
	for _, patient := range patients {
		patientResp := NursePatient{
			NurseID:              patient.NurseID,
			NurseFirstName:       patient.NurseFirstName,
			NurseLastName:        patient.NurseLastName,
			PatientID:            patient.PatientID,
			PatientFirstName:     patient.PatientFirstName,
			PatientLastName:      patient.PatientLastName,
			Age:                  patient.Age,
			Sex:                  patient.Sex,
			BloodType:            patient.BloodType,
			PhoneNumber:          patient.PhoneNumber,
			Address:              patient.Address,
			DOB:                  patient.DOB,
			AssignedDoctorID:     patient.AssignedDoctorID,
			BodyTempertature:     patient.BodyTempertature,
			PulseRate:            patient.PulseRate,
			RespirationRate:      patient.RespirationRate,
			SystolicPressure:     patient.SystolicPressure,
			DiastolicPressure:    patient.DiastolicPressure,
			CurrentPrescribedMed: patient.CurrentPrescribedMed,
			CurrentDisease:       patient.CurrentDisease,
		}
		resp.Patients = append(resp.Patients, patientResp)
	}

	ctx.JSON(200, resp)
}

type DoctorDashboardResp struct {
	Patients []DoctorPatient `json:"patients"`
}
type DoctorPatient struct {
	ID                   int       `json:"patient_id"`
	FirstName            string    `json:"first_name"`
	LastName             string    `json:"last_name"`
	Age                  int       `json:"age"`
	Sex                  string    `json:"sex"`
	BloodType            string    `json:"blood_type"`
	PhoneNumber          string    `json:"phone_number"`
	Address              string    `json:"address"`
	DOB                  time.Time `json:"dob"`
	AssignedDoctorID     int       `json:"assigned_doctor_id"`
	BodyTempertature     float64   `json:"body_temperature"`
	PulseRate            int       `json:"pulse_rate"`
	RespirationRate      int       `json:"respiration_rate"`
	SystolicPressure     int       `json:"systolic_pressure"`
	DiastolicPressure    int       `json:"diastolic_pressure"`
	CurrentPrescribedMed string    `json:"current_prescribed_med"`
	CurrentDisease       string    `json:"current_disease"`
}

func (h *DashboardHandler) GetDoctorDashboard(ctx *gin.Context) {
	didStr := ctx.Query("doctor_id")
	if didStr == "" {
		ctx.JSON(400, gin.H{"error": "doctor_id is required"})
		return
	}
	did, err := strconv.Atoi(didStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "doctor_id must be integer"})
		return
	}
	patients, err := h.repo.SelectDoctorDashboard(did)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var resp DoctorDashboardResp
	for _, patient := range patients {
		patientResp := DoctorPatient{
			ID:                   patient.ID,
			FirstName:            patient.FirstName,
			LastName:             patient.LastName,
			Age:                  patient.Age,
			Sex:                  patient.Sex,
			BloodType:            patient.BloodType,
			PhoneNumber:          patient.PhoneNumber,
			Address:              patient.Address,
			DOB:                  patient.DOB,
			AssignedDoctorID:     patient.AssignedDoctorID,
			BodyTempertature:     patient.BodyTempertature,
			PulseRate:            patient.PulseRate,
			RespirationRate:      patient.RespirationRate,
			SystolicPressure:     patient.SystolicPressure,
			DiastolicPressure:    patient.DiastolicPressure,
			CurrentPrescribedMed: patient.CurrentPrescribedMed,
			CurrentDisease:       patient.CurrentDisease,
		}
		resp.Patients = append(resp.Patients, patientResp)
	}

	ctx.JSON(200, resp)
}

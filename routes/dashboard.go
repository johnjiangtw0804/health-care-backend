package routes

import (
	repository "health-care-backend/repository"
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
	dashboard, err := h.repo.GetPatientDashboard()
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

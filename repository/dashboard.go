package repository

import (
	"fmt"
	model "health-care-backend/repository/model"
)

type Dashboard interface {
	GetPatientDashboard() (model.PatientDashboardView, error)
	GetDoctorDashboard()
	GetNurseDashboard()
}

type dashboardRepo struct {
	db *GormDatabase
}

func NewDashboardRepo(db *GormDatabase) Dashboard {
	return &dashboardRepo{db: db}
}

func (d *dashboardRepo) GetPatientDashboard() (model.PatientDashboardView, error) {
	var record model.PatientDashboardView
	if err := d.db.DB.Raw(`SELECT * FROM public.patient_dashboard_view`).Scan(&record).Error; err != nil {
		return model.PatientDashboardView{}, err
	}

	fmt.Println(record)
	return record, nil
}

func (d *dashboardRepo) GetDoctorDashboard() {
}

func (d *dashboardRepo) GetNurseDashboard() {

}

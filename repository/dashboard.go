package repository

import (
	"fmt"
	model "health-care-backend/repository/model"
)

type Dashboard interface {
	SelectPatientDashboard(int) ([]model.PatientDashboardView, error)
	SelectDoctorDashboard(did int) ([]model.DoctorDashboardView, error)
	SelectNurseDashboard(nid int) ([]model.NurseDashboardView, error)
}

type dashboardRepo struct {
	db *GormDatabase
}

func NewDashboardRepo(db *GormDatabase) Dashboard {
	return &dashboardRepo{db: db}
}

func (d *dashboardRepo) SelectPatientDashboard(pid int) ([]model.PatientDashboardView, error) {
	var records []model.PatientDashboardView
	if err := d.db.DB.Raw(`SELECT * FROM public.patient_dashboard_view WHERE id = ?`, pid).Scan(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (d *dashboardRepo) SelectDoctorDashboard(did int) ([]model.DoctorDashboardView, error) {
	var records []model.DoctorDashboardView
	if err := d.db.DB.Raw(`SELECT * FROM public.doctor_dashboard_view WHERE assigned_doctor_id = ?`, did).Scan(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (d *dashboardRepo) SelectNurseDashboard(nid int) ([]model.NurseDashboardView, error) {
	var records []model.NurseDashboardView
	if err := d.db.DB.Raw(`SELECT * FROM public.nurse_dashboard_view WHERE nurse_id = ?`, nid).Scan(&records).Error; err != nil {
		return nil, err
	}
	fmt.Println(records)
	return records, nil
}

package repository

type Dashboard interface {
	GetPatientDashboard()
	GetDoctorDashboard()
	GetNurseDashboard()
}

type dashboardRepo struct {
	db *GormDatabase
}

func NewDashboardRepo(db *GormDatabase) Dashboard {
	return &dashboardRepo{db: db}
}

func (r *dashboardRepo) GetPatientDashboard() {

}

func (r *dashboardRepo) GetDoctorDashboard() {
}

func (r *dashboardRepo) GetNurseDashboard() {

}

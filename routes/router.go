package routes

import (
	envconfig "health-care-backend/envconfig"
	"health-care-backend/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Register(
	router *gin.Engine,
	logger *zap.Logger,
	db *repository.GormDatabase,
	env *envconfig.Env,
) *gin.Engine {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(config))

	dashboardRepo := repository.NewDashboardRepo(db)

	dashboardHandler := NewDashboardHandler(logger, dashboardRepo)

	router.GET("/api/dashboard/patient", dashboardHandler.GetPatientDashboard)
	router.GET("/api/dashboard/doctor", dashboardHandler.GetDoctorDashboard)
	router.GET("/api/dashboard/nurse", dashboardHandler.GetNurseDashboard)
	return router
}

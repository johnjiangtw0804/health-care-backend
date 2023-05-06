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

	// dashboardRepo := repository.DashboardRepo(db)

	// noAuthRouters := router.Group("")
	// userHandler := NewUserHandler(logger, userRepo, api.NewMapUtilities(env.GOOGLE_MAP_API_KEY), api.NewEventsSearcher(env.TICKET_MASTER_API_KEY))

	// noAuthRouters.GET("/api/user/events", userHandler.listEvents)
	return router
}

package router

import (
	"database/sql"
	"meeting-room-booking/internal/controller"
	"meeting-room-booking/internal/middleware"
	"meeting-room-booking/internal/repository"
	"meeting-room-booking/internal/service"
	"time"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Главный router
func Router(r *gin.Engine, db *sql.DB) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Создание middleware для аутентификации
	repo := repository.NewUserRepository(db)
	authMiddleware := middleware.BasicAuthMiddleware(repo)

	// Группа для версий API
	V1 := r.Group("v1")
	{
		// Открытые маршруты, не требующие аутентификации
		V1.POST("/register", registerHandler(db))
		V1.POST("/login", loginHandler(db))
		V1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		// Защищенные маршруты, к которым применяется authMiddleware
		protected := V1.Group("/")
		protected.Use(authMiddleware)
		{
			UserRouter(protected, db)
			BookingRouter(protected, db)
		}
	}

	r.Static("/static", "./web/booking-frontend/build/static")
	r.StaticFile("/favicon.ico", "./web/booking-frontend/build/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.File("./web/booking-frontend/build/index.html")
	})

	r.NoRoute(func(c *gin.Context) {
		c.File("./web/booking-frontend/build/index.html")
	})
}

// UserRouter — маршруты для работы с пользователями
func UserRouter(router *gin.RouterGroup, db *sql.DB) {
	repo := repository.NewUserRepository(db)
	service := service.NewUserService(repo)
	controller := controller.NewUserController(service)

	router.GET("/users", controller.GetAll)
	router.GET("/user/:id", controller.GetByID)
	router.GET("/user/username/:username", controller.GetByUsername)
	router.POST("/user", controller.Create)
	router.PUT("/user/:id", controller.Update)
	router.DELETE("/user/:id", controller.Delete)
}

// BookingRouter — маршруты для работы с бронированием
func BookingRouter(router *gin.RouterGroup, db *sql.DB) {
	repo := repository.NewBookingRepository(db)
	service := service.NewBookingService(repo)
	controller := controller.NewBookingController(service)

	router.GET("/bookings", controller.GetAll)
	router.GET("/booking/:id", controller.GetByID)
	router.POST("/booking", controller.Create)
	router.PUT("/booking/:id", controller.Update)
	router.DELETE("/booking/:id", controller.Delete)
}

// registerHandler — обработчик для регистрации
func registerHandler(db *sql.DB) gin.HandlerFunc {
	repo := repository.NewUserRepository(db)
	service := service.NewUserService(repo)
	controller := controller.NewUserController(service)
	return controller.Register
}

// loginHandler — обработчик для входа
func loginHandler(db *sql.DB) gin.HandlerFunc {
	repo := repository.NewUserRepository(db)
	service := service.NewUserService(repo)
	controller := controller.NewUserController(service)
	return controller.Login
}

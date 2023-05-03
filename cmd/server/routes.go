package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/ZeinaMFI/internal/api"
	"github.com/michaelgbenle/ZeinaMFI/internal/middleware"
	"github.com/michaelgbenle/ZeinaMFI/internal/ports"
	"time"
)

//SetupRouter is where router endpoints are called
func SetupRouter(handler *api.HTTPHandler, repository ports.Repository) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r := router.Group("/")
	{
		r.GET("/live", handler.LiveHandler)
		r.GET("/users", handler.GetAllUsersHandler)
		r.POST("/register", handler.SignUpHandler)
		r.POST("/login", handler.LoginHandler)
	}

	// authorizeAdmin authorizes all authorized users handlers
	authorizeAdmin := r.Group("/admin")
	authorizeAdmin.Use(middleware.AuthorizeAdmin(repository.FindUserByEmail, repository.TokenInBlacklist))
	{
		authorizeAdmin.PATCH("deposit", handler.DepositHandler)
		authorizeAdmin.PATCH("withdraw", handler.WithdrawHandler)
		authorizeAdmin.PATCH("lock", handler.LockSavingsHandler)
		authorizeAdmin.GET("transactions/:account_no", handler.TransactionsHandler)
	}

	return router
}

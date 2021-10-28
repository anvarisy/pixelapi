package controllers

import (
	"net/http"

	"github.com/anvarisy/pixelapi/docs"
	"github.com/anvarisy/pixelapi/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) initializeRoutes() {
	docs.SwaggerInfo.Title = "Backend Pixel "
	docs.SwaggerInfo.Description = "Backend untuk kebutuhan Pixel"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	s.Router.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	s.Enforcer.AddPolicy("Admin", "Admin", "read")
	s.Enforcer.AddPolicy("Admin", "Admin", "write")
	s.Enforcer.AddPolicy("Buyer", "Buyer", "write")
	s.Enforcer.AddPolicy("Buyer", "Buyer", "read")
	api := s.Router.Group("/")
	{
		api.GET("/", func(c *gin.Context) {
			c.Redirect(http.StatusFound, "/doc/index.html")
		})
		api.POST("/login", s.LoginController)
		api.POST("/register", s.RegisterController)
		api.POST("/create-user", middlewares.TokenAuthentication(), middlewares.Authorize("Admin", "write", s.Enforcer), s.CreateUserController)
		// Stuff
		api.GET("/stuff", middlewares.TokenAuthentication(), middlewares.Authorize("Admin", "read", s.Enforcer), s.GetAllStuffController)
		api.POST("/stuff", middlewares.TokenAuthentication(), middlewares.Authorize("Admin", "write", s.Enforcer), s.CreateStuffController)
		api.POST("/stuff/update/:id", middlewares.TokenAuthentication(), middlewares.Authorize("Admin", "write", s.Enforcer), s.UpdateStuffController)
		api.POST("/stuff/delete/multiple", middlewares.TokenAuthentication(), middlewares.Authorize("Admin", "write", s.Enforcer), s.DeleteMultipleStuffController)
		api.GET("/stuff/:id", middlewares.TokenAuthentication(), middlewares.Authorize("Admin", "read", s.Enforcer), s.GetStuffByIDController)
		api.POST("/stuff/delete/:id", middlewares.TokenAuthentication(), middlewares.Authorize("Admin", "write", s.Enforcer), s.DeleteStuffController)
		api.GET("/stuff/cosumer", s.GetStuffByCostumerController)
		// Transaction
		api.POST("/transaction", middlewares.TokenAuthentication(), s.CreateTransactionController)

	}
}

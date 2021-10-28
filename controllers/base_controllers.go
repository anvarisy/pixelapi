package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anvarisy/pixelapi/middlewares"
	"github.com/anvarisy/pixelapi/models"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	DB       *gorm.DB
	Router   *gin.Engine
	Enforcer *casbin.Enforcer
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	server.DB, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to %s database mysql", "")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", "")
	}

	server.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.User{},
		&models.Stuff{},
		&models.Transaction{},
		&models.TransactionDetail{},
	)

	err_seed := server.DB.Where("username = ?", "koala").Take(&models.User{}).Error
	if err_seed == nil {
		log.Println("Username already insert")
	}
	server.DB.Create(&models.User{Username: "koala", UserPassword: "panda",
		UserFullname: "Koala Panda", UserMobile: "6285219529352", IsAdmin: true})

	server.Router = gin.Default()
	server.Router.Use(middlewares.CORSMiddleware())
	adapter, err := gormadapter.NewAdapterByDB(server.DB)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize casbin adapter: %v", err))
	}

	// Load model configuration file and policy store adapter
	enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	if err != nil {
		panic(fmt.Sprintf("failed to create casbin enforcer: %v", err))
	}
	server.Enforcer = enforcer
	server.initializeRoutes()
	server.Enforcer.AddGroupingPolicy("koala", "Admin")

}
func (server *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/anvarisy/pixelapi/auths"
	"github.com/anvarisy/pixelapi/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginController ... User Login
// @Summary User Login
// @Description Url API untuk login setiap user
// @Tags Soal nomor 2
// @Accept  json
// @Produce  json
// @Param User body models.UserLogin true "User Data"
// @Success 200 {object} models.UserLoginSuccess
// @Router /login [post]
func (s *Server) LoginController(c *gin.Context) {
	errList := map[string]string{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":      http.StatusUnprocessableEntity,
			"first error": errList,
		})
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		errList["Unmarshal_error"] = err.Error()
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	res, err := s.Signin(user.Username, user.UserPassword)
	if err != nil {
		errList["error_found"] = err.Error()
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": res,
	})

}

func (s *Server) Signin(username, password string) (map[string]interface{}, error) {
	var err error
	data := make(map[string]interface{})
	u := models.User{}
	err = s.DB.Where("username = ?", username).Take(&u).Error
	if err != nil {
		fmt.Println("this is the error getting the user: ", err)
		return nil, err
	}
	err = auths.VerifyPassword(u.UserPassword, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		fmt.Println("this is the error hashing the password: ", err)
		return nil, err
	}
	token, _ := auths.CreateToken(u.Username, u.IsAdmin)
	data["username"] = u.Username
	data["user_fullname"] = u.UserFullname
	data["user_mobile"] = u.UserMobile
	data["user_token"] = token
	return data, nil
}

// CreateUserController ... Create User
// @Summary User Create
// @Description Digunakan untuk admin yang ingin menambahkan user atau admin baru
// @Tags Soal Nomor 1
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param User body models.User true "Add User"
// @Success 201 {object} models.UserRegisterSuccess
// @Router /create-user [post]
func (s *Server) CreateUserController(c *gin.Context) {
	errList := map[string]string{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":      http.StatusUnprocessableEntity,
			"first error": errList,
		})
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		errList["Unmarshal_error"] = err.Error()
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	res, err := user.CreateUser(s.DB)
	if err != nil {
		errList["register_failed"] = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	if res.IsAdmin {
		s.Enforcer.AddGroupingPolicy(res.Username, "Admin")
	} else {
		s.Enforcer.AddGroupingPolicy(res.Username, "Buyer")
	}
	res.UserPassword = ""
	c.JSON(http.StatusCreated, gin.H{
		"status":   http.StatusCreated,
		"response": res,
	})

}

// RegisterController ... Consumer Register
// @Summary Consumer Register
// @Description Digunakan untuk user yang akan melakukan registrasi
// @Tags Soal Nomor 1
// @Accept  json
// @Produce  json
// @Param User body models.UserRegister true "Add User"
// @Success 201 {object} models.UserRegisterSuccess
// @Router /register [post]
func (s *Server) RegisterController(c *gin.Context) {
	errList := map[string]string{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":      http.StatusUnprocessableEntity,
			"first error": errList,
		})
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		errList["Unmarshal_error"] = err.Error()
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	res, err := user.CreateUser(s.DB)
	if err != nil {
		errList["register_failed"] = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	s.Enforcer.AddGroupingPolicy(res.Username, "Buyer")
	res.UserPassword = ""
	c.JSON(http.StatusCreated, gin.H{
		"status":   http.StatusCreated,
		"response": res,
	})
}

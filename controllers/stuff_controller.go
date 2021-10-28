package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/anvarisy/pixelapi/models"
	"github.com/gin-gonic/gin"
)

// GetAllStuffController ... List Stuffes
// @Summary Stuffes Get
// @Description Mengambil semua data barang
// @Tags Soal nomor 3
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {array} models.Stuff
// @Router /stuff [get]
func (s *Server) GetAllStuffController(c *gin.Context) {
	stuff := models.Stuff{}
	res, err := stuff.GetAllStuff(s.DB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": res,
	})
}

// CreateStuffController ... Stuff Create
// @Summary Stuff Create
// @Description membuat data barang
// @Tags Soal nomor 3
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param Stuff body models.StuffSerializer true "Stuff Data"
// @Success 201 {object} models.Stuff
// @Router /stuff [post]
func (s *Server) CreateStuffController(c *gin.Context) {
	stuff := models.Stuff{}
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
	err = json.Unmarshal(body, &stuff)
	if err != nil {
		errList["Unmarshal_error"] = err.Error()
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	res, err := stuff.CreateStuff(s.DB)
	if err != nil {
		errList["input_failed"] = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":   http.StatusCreated,
		"response": res,
	})
}

// UpdateStuffController ... Stuff Update
// @Summary Stuff Update
// @Description API URL untuk update barang
// @Tags Soal nomor 3
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id path string true "id"
// @Param Stuff body models.StuffSerializer true "Stuff Data"
// @Success 202 {object} models.Stuff
// @Router /stuff/update/{id} [post]
func (s *Server) UpdateStuffController(c *gin.Context) {
	stuff := models.Stuff{}
	id := c.Param("id")
	errList := map[string]string{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	err = json.Unmarshal(body, &stuff)
	if err != nil {
		errList["Unmarshal_error"] = err.Error()
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	// log.Println(&cctv)
	err = s.DB.Model(&stuff).Where("id = ?", id).Updates(map[string]interface{}{"stuff_name": stuff.StuffName, "stuff_price": stuff.StuffPrice, "stuff_stock": stuff.StuffStock}).Error
	if err != nil {
		errList["Internal_Error"] = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":   http.StatusAccepted,
		"response": &stuff,
	})
}

// DeleteMultipleStuffController ... Delete Multiple Stuff
// @Summary Stuff Multiple Delete
// @Description API URL untuk menghapus beberapa barang
// @Tags Soal nomor 3
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param Stuff ID body models.StuffID true "Array of stuff ID"
// @Success 202
// @Router /stuff/delete/multiple [post]
func (server *Server) DeleteMultipleStuffController(c *gin.Context) {
	ids := models.StuffID{}
	stuff := models.Stuff{}
	errList := map[string]string{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	err = json.Unmarshal(body, &ids)
	if err != nil {
		errList["Unmarshal_error"] = "Cannot unmarshal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	for _, item := range ids.ID {
		err = server.DB.Where("id = ?", item).Delete(&stuff).Error
		if err != nil {
			errList["Internal_Error"] = "Process stoped cause failed to delete cctv " + item
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusInternalServerError,
				"error":  errList,
			})
			return
		}
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":   http.StatusAccepted,
		"response": "Delete Complete",
	})
}

// GetStuffByIDController ... Stuff Get By ID
// @Summary Stuff Get By ID
// @Description API URL untuk mengambil data barang berdasarkan id
// @Tags Soal nomor 3
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id path string true "id cctv"
// @Success 200 {object} models.Stuff
// @Router /stuff/{id} [get]
func (server *Server) GetStuffByIDController(c *gin.Context) {
	stuff := models.Stuff{}
	id := c.Param("id")
	err := server.DB.Where("id = ?", id).Take(&stuff).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": stuff,
	})
}

// DeleteStuffController ... Stuff Delete
// @Summary Stuff Delete
// @Description API URL untuk menghapus barang
// @Tags Soal nomor 3
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id path string true "id"
// @Success 202
// @Router /stuff/delete/{id} [post]
func (server *Server) DeleteStuffController(c *gin.Context) {
	stuff := models.Stuff{}
	id := c.Param("id")
	err := server.DB.Where("id = ?", id).Delete(&stuff).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err,
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":   http.StatusAccepted,
		"response": "Stuff has been deleted",
	})
}

// GetStuffByCostumerController ... -
// @Summary Stuff for consumer
// @Description APi barang untuk consumer
// @Tags Soal nomor 4
// @Accept  json
// @Produce  json
// @Param page query int false "used for page"
// @Param page_size query int false "used for page size"
// @Success 200 {array} models.Stuff
// @Router /stuff/cosumer [get]
func (s *Server) GetStuffByCostumerController(c *gin.Context) {
	stuffes := []models.Stuff{}

	err := s.DB.Scopes(Paginate(c)).Find(&stuffes).Error
	if err != nil {
		c.JSON(http.StatusNoContent, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": stuffes,
	})
}

package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/anvarisy/pixelapi/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

// CreateTransactionController ... Stuff Create
// @Summary Transaction Create
// @Description membuat data transaksi
// @Tags Soal no 5
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param Stuff body models.TransactionSerializer true "Transaction Data"
// @Success 201 {object} models.Transaction
// @Router /transaction [post]
func (s *Server) CreateTransactionController(c *gin.Context) {
	t := models.TransactionSerializer{}
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
	err = json.Unmarshal(body, &t)
	if err != nil {
		errList["Unmarshal_error"] = err.Error()
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	var amount int64 = 0
	for _, i := range t.Detail {
		status := s.CheckStock(int64(i.StuffID), i.Count)
		if !status {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"status": http.StatusUnprocessableEntity,
				"error":  "Barang dengan id " + strconv.Itoa(i.StuffID) + " tidak memenuhi permintaan",
			})
			return
		}
		amount += (i.Count * s.GetPrice(i.StuffID))

	}

	// Check stock

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": "AMan",
	})
}

func (s *Server) CheckStock(id_stuff, order int64) bool {
	stuff := models.Stuff{}
	ts := []models.Transaction{}
	err := s.DB.Where("id = ?", id_stuff).Take(&stuff).Error
	if err != nil {
		fmt.Println("this is the error getting the stuff: ", err)
		return false
	}
	if (stuff.StuffStock - order) < 0 {
		fmt.Println("stock not enough for this stuff: ", stuff.StuffName)
		return false
	} else {
		s.DB.Where("is_paid = ?", false).Preload(clause.Associations).Find(&ts)
		var ordered int64 = 0
		for _, item_transaction := range ts {
			for _, item := range item_transaction.Detail {
				ordered += item.Count
			}
		}
		log.Println(ordered)
		if (stuff.StuffStock - order - ordered) < 0 {
			fmt.Println("stock not enough for this stuff: ", stuff.StuffName)
			return false
		} else {
			return true
		}

	}
}

func (s *Server) GetPrice(id int) int64 {
	stuff := models.Stuff{}
	s.DB.Where("id = ?", id).Take(&stuff)
	return stuff.StuffPrice
}

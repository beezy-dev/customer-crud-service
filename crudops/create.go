package crudops

import (
	"log"

	"github.com/beezy-dev/traveler-crud-service/inits"
	"github.com/beezy-dev/traveler-crud-service/models"
	"github.com/gin-gonic/gin"
)

func CreateOps(c *gin.Context) {

	// data from API request body

	var body struct {
		Uuid       string
		Firstmame  string
		Lastname   string
		Multihop   bool
		Pprocessed bool
		Wprocessed string
		Reprocess  bool
		Offloaded  bool
	}

	c.Bind(&body)

	// process data to database

	create := models.Traveler{
		Uuid:       body.Uuid,
		Firstname:  body.Firstmame,
		Lastname:   body.Lastname,
		Multihop:   body.Multihop,
		Pprocessed: body.Pprocessed,
		Wprocessed: body.Wprocessed,
		Reprocess:  body.Reprocess,
		Offloaded:  body.Offloaded,
	}

	r := inits.DB.Create(&create)

	if r.Error != nil {
		c.JSON(400, gin.H)
		log.Println("Error: C operation failed with error")
		log.Fatal(r.Error)
	} else {
		c.JSON(200, gin.H{
			""
		})
	}

}

// Uuid       string
// Firstmame  string
// Lastname   string
// Multihop   bool
// Pprocessed bool
// Wprocessed string
// Reprocess  bool
// Offloaded  bool

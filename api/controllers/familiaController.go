package controllers

import (
	"errors"
	// "gorm-test/database"
	// "gorm-test/models"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/wilmer88/thefamila/api/database"
	"github.com/wilmer88/thefamila/api/models"
	"gorm.io/gorm"
)

type Familia struct {
	Db *gorm.DB
}

func New() *Familia {
	db := database.InitDb()
	db.AutoMigrate(&models.Fammember{})
	return &Familia{Db: db}
}

//create user
func (repository *Familia) CreateUser(c *gin.Context) {
	var member models.Fammember
	c.BindJSON(&member)
	err := models.CreateUser(repository.Db, &member)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, member)
}

//get users
func (repository *Familia) GetUsers(c *gin.Context) {
	var member []models.Fammember
	err := models.GetUsers(repository.Db, &member)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, member)
}

//get user by id
func (repository *Familia) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var member models.Fammember
	err := models.GetUser(repository.Db, &member, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, member)
}

// update user
func (repository *Familia) UpdateUser(c *gin.Context) {
	var member models.Fammember
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.GetUser(repository.Db, &member, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&member)
	err = models.UpdateUser(repository.Db, &member)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, member)
}

// delete user
func (repository *Familia) DeleteUser(c *gin.Context) {
	var member models.Fammember
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteUser(repository.Db, &member, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
package controller

import (
	"github.com/gin-gonic/gin"
	"api-auth-test/model"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}
	user, err := model.Model.CreateUser(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, user)
}

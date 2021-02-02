package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"api-auth-test/auth"
	"api-auth-test/model"
	"api-auth-test/service"
	"net/http"
)

func Login(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	user, err := model.Model.GetUserByEmail(u.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	authData, err := model.Model.CreateAuth(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	var authD auth.Details
	authD.UserId = authData.UserID
	authD.AuthUid = authData.AuthUUID

	token, loginErr := service.Authorize.SignIn(authD)
	if loginErr != nil {
		c.JSON(http.StatusForbidden, "Please try login later")
		return
	}
	c.JSON(http.StatusOK, token)
}

func LogOut(c *gin.Context) {
	au, err := auth.ExtractTokenAuth(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	delErr := model.Model.DeleteAuth(au)
	if delErr != nil {
		log.Println(delErr)
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}

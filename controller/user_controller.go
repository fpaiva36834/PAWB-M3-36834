package controller

import (
	"awesomeProject/entity"
	"awesomeProject/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "select all users",
		"users":   service.GetAllUsers(),
	})
}

func Register(c *gin.Context) {
	var user entity.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
		})
		return
	}
	user = service.InsertUser(user)
	c.JSON(200, gin.H{
		"message": "add user",
		"user":    user,
	})
}

func Profile(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "id error",
		})
		return
	}
	user, err := service.GetUserByID(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "select user by id",
		"user":    user,
	})
}

func UpdateProfile(c *gin.Context) {
	userId, iderr := strconv.ParseUint(c.Param("id"), 10, 64)
	if iderr != nil {
		c.JSON(400, gin.H{
			"message": "id error",
		})
		return
	}
	var user entity.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
		})
		return
	}
	user.ID = userId
	user, err = service.UpdateUserByID(user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "update user by id",
		"user":    user,
	})
}

func DeleteAccount(c *gin.Context) {
	userId, iderr := strconv.ParseUint(c.Param("id"), 10, 64)
	if iderr != nil {
		c.JSON(400, gin.H{
			"message": "id error",
		})
		return
	}
	service.DeleteUserByID(userId)
	c.JSON(200, gin.H{
		"message": "delete user by id",
	})
}

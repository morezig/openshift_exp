package controllers

import (
	"fmt"
	"net/http"
	"openshift_exp/oc-integrator/go-api/models"
	"time"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func LoginJWT(username string, password string, ctx *gin.Context) (interface{}, bool) {
	fmt.Println("username:", username, " ,password:", password)
	if username == "admin" && password == "admin" {
		return &models.User{Username: "admin"}, true
	} else if userID, OK := UserAuthen(username, password, ctx); OK {
		return &models.User{Username: userID}, true
	} else {
		return nil, false
	}
}

func LoginResponse(ctx *gin.Context, code int, token string, expire time.Time) {
	dataMap := &Resp{
		Code: http.StatusOK,
		Data: map[string]interface{}{
			"token": token,
		},
	}
	ctx.JSON(http.StatusOK, dataMap)
}

func GetUserInfo(ctx *gin.Context) {
	dataMap := &Resp{
		Code: http.StatusOK,
		Data: map[string]interface{}{
			"roles": []string{"admin"},
			"name":  "admin",
		},
	}
	ctx.JSON(http.StatusOK, dataMap)
}

func LogoutHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": "success",
	})
}

func GetUsers(ctx *gin.Context) {

}

func Login(ctx *gin.Context) {
	var loginJSON models.User
	ctx.BindJSON(&loginJSON)
	username := loginJSON.Username
	password := loginJSON.Password
	// fmt.Printf("username :%s, password :%s\n", username, password)
	if username == "admin" && password == "admin" {
		fmt.Println("admin login")
		ctx.JSON(http.StatusOK, gin.H{
			"user":  loginJSON,
			"token": "fake token",
		})
	} else if userID, OK := UserAuthen(username, password, ctx); OK {
		fmt.Println("userId:", userID)
		ctx.JSON(http.StatusOK, gin.H{
			"user":  loginJSON,
			"token": "fake token",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "Username or password is incorrect",
		})
	}
}

func UserAuthen(userID string, password string, c *gin.Context) (string, bool) {
	if len(userID) > 0 && len(password) > 0 {
		return models.CheckAuth(userID, password)
	}
	return userID, false
}

func UserAuthor(user interface{}, c *gin.Context) bool {
	if v, ok := user.(string); ok && v == "admin" {
		return true
	}
	return false
}

func UnAuthor(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

func UserCreate(c *gin.Context) {
	var usrobj models.User
	err := c.BindJSON(&usrobj)
	if err != nil {
		fmt.Println("Error:", err)
	}
	usrobj.Insert()
	c.JSON(http.StatusOK, usrobj)
}

func UserUpdate(c *gin.Context) {
	// usrform := make(map[string]interface{})
	var usrobj models.User
	err := c.BindJSON(&usrobj)
	if err != nil {
		fmt.Println("Error:", err)
	}
	usrobj.Update()
	c.JSON(http.StatusOK, usrobj)
}

func UserDelete(c *gin.Context) {
	id := c.Param("id")
	models.DelUserByID(id)
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// func UserAction(c *gin.Context) {
// 	action := c.Param("action")
// 	switch action {
// 	case "form":
// 		GetModelForm(c, "user")
// 	case "view":
// 		GetModelView(c, "user")
// 	case "grid":
// 		GetModelGrid(c, "user")
// 	case "options":
// 		GetUserOptions(c)
// 	case "undefined":
// 		c.JSON(http.StatusOK, gin.H{
// 			"Code":    http.StatusNotFound,
// 			"Message": "Record Not Found",
// 		})
// 	default:
// 		UserQueryByID(c, action)
// 	}
// }

// func UserQueryByID(c *gin.Context, id string) {
// 	user, _ := models.GetUserByID(id)
// 	user.FillUserRequiredIds()
// 	c.JSON(http.StatusOK, user)
// }

// func CreateUser(c *gin.Context) {
// 	var userObj models.User
// 	c.BindJSON(&userObj)
// 	userObj.Insert()
// }

// func UpdateeUser(c *gin.Context) {
// 	var userObj models.User
// 	c.BindJSON(&userObj)
// 	userObj.Update()
// }

// func CreateTmpUser(c *gin.Context) {
// 	user := &models.User{
// 		Username: "admin",
// 		Email:    "admin@cenoq.com",
// 		Password: "admin",
// 	}
// 	user.Password = common.Md5(user.Email + user.Password)
// 	err := user.Insert()
// 	if err == nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"succeed": true,
// 		})
// 	}
// }

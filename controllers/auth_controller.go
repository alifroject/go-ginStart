package controllers

import (
	"gin-quickstart/models"
	"gin-quickstart/services"
	"gin-quickstart/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignupInput struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AdminLoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func SignUp(c *gin.Context) {
	var input SignupInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	user := models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
	}

	createdUser, err := services.SignUp(&user)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusCreated, createdUser)
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	token, user, err := services.Login(input.Email, input.Password)
	if err != nil {
		utils.RespondError(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})

}

func AdminLogin(c *gin.Context) {
    var input AdminLoginInput

    if err := c.ShouldBindJSON(&input); err != nil {
        utils.RespondError(c, http.StatusBadRequest, err.Error())
        return
    }

    token, user, err := services.AdminLogin(input.Email, input.Password)
    if err != nil {
        utils.RespondError(c, http.StatusUnauthorized, err.Error())
        return
    }

    // Enforce Admin role
    if user.Role != "admin" {
        utils.RespondError(c, http.StatusForbidden, "Access denied: Admins only")
        return
    }

    // Cookie settings
    c.SetSameSite(http.SameSiteStrictMode)

    c.SetCookie(
        "token",
        token,
        3600*24,
        "/",
        "",
        true,  // secure
        true,  // httpOnly
    )

    utils.RespondJSON(c, http.StatusOK, gin.H{
        "user": user,
    })
}



func GetMe(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	c.JSON(http.StatusOK, user)
}

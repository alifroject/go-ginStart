package controllers

import (
    "net/http"

    "gin-quickstart/models"
    "gin-quickstart/services"
    "gin-quickstart/utils"

    "github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        utils.RespondError(c, http.StatusBadRequest, "Invalid JSON")
        return
    }

    if err := services.CreateUser(&user); err != nil {
        utils.RespondError(c, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondJSON(c, http.StatusCreated, user)
}

func GetUsers(c *gin.Context) {
    users, err := services.GetUsers()
    if err != nil {
        utils.RespondError(c, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondJSON(c, http.StatusOK, users)
}



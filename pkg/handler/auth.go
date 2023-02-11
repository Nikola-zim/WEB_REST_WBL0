package handler

import (
	"WEB_REST_exm0302/static"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	//структура в которой будем записывать данные из json от пользователей
	var input static.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body") //StatusBadRequest = 400 "некоректные данные"
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) //Код 500 - ошибка сервера
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{ //Код 200
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	//структура в которой будем записывать данные из json от пользователей
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()) //StatusBadRequest = 400 "некоректные данные"
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) //Код 500 - ошибка сервера
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{ //Код 200
		"token": token,
	})
}

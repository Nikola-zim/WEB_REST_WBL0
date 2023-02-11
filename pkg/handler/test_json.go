package handler

import (
	"WEB_REST_exm0302/static"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) readId(c *gin.Context) {
	c.HTML(http.StatusOK, "test_json.html", gin.H{
		"block_title": "Test page",
	})
}

func (h *Handler) showJson(c *gin.Context) {
	desiredJsonId := c.PostForm("json_id")
	if desiredJsonId == "" {
		c.HTML(http.StatusOK, "test_json.html", gin.H{
			"errors_block": "Ошибка! Пожалуйста, введите id",
		})
	}

	//var myid uint64
	//myid, _ = strconv.ParseUint(desiredJsonId, 0, 0)
	myTestMap, err := h.services.JsonRW.ReadFromCash(desiredJsonId)
	if err != nil {
		c.HTML(http.StatusOK, "test_json.html", gin.H{
			"errors_block": "Ошибка! Такого id не существует",
		})
	}
	myTestMap2, _ := json.Marshal(myTestMap)
	c.HTML(http.StatusOK, "test_json.html", gin.H{
		"block_output": string(myTestMap2),
	})
}

// Функция для тестирования ввода json через http запрос
func (h *Handler) writeJson(c *gin.Context) {
	//структура в которой будем записывать данные из json
	var inputJson static.Json

	//Проверка на нужный формат Json
	if err := c.BindJSON(&inputJson); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body") //StatusBadRequest = 400 "некоректные данные"
		return
	}
	//Сохранение в кеш
	err := h.services.JsonRW.WriteInCash(inputJson)
	//Cохранение в БД
	errDB := h.services.JsonRW.WriteInDB(inputJson)
	//Проверка что небыло ошибок и при записи в БД и при записи в кэш
	if err != nil || errDB != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) //Код 500 - ошибка сервера
		return
	}
}

package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//test := cash.NewTestCash()
//test.Testmap[0] = "HIHIHIHIHI"

func (h *Handler) showTestHome(c *gin.Context) {
	c.HTML(http.StatusOK, "test.html", gin.H{
		"block_title": "Test page",
	})
}

func (h *Handler) showResultTestHome(c *gin.Context) {

	//myNum := c.PostForm("json_id")
	//logrus.Debugf(myNum)
	//c.HTML(http.StatusOK, "test.html", gin.H{
	//	"block_output": myNum,
	//})
	myNum := c.PostForm("json_id")
	if myNum == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body") //StatusBadRequest = 400 "некоректные данные"
		return
	}

	err := h.services.CashNumbers.AppendNumberInCash(myNum)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) //Код 500 - ошибка сервера
		return
	}

	myTestMap, err := h.services.CashNumbers.ReadNumbersFromCash()
	c.HTML(http.StatusOK, "test.html", gin.H{
		"block_output": myTestMap,
	})
}

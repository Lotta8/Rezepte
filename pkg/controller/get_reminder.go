package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type responseReminder struct {
	Id   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func (h Handler) GetReminder(context *gin.Context) {
	param := context.Param("id") // konventieren in eine Zahl
	id, err := strconv.Atoi(param)
	if err != nil { //Prüfung falls nicht konvertiert, bricht ab
		context.JSON(http.StatusInternalServerError, "invalid id")
		return
	}
	reminder, err := h.repository.Get(id) // Wieder den Handler auf welchem ich das Repository aufrufen kann. Get ist auf dem Repository definiert.
	if err != nil {
		_ = context.AbortWithError(http.StatusBadRequest, err)
	}
	context.JSON(http.StatusOK, responseReminder{Id: reminder.Id, Task: reminder.Task, Done: reminder.Done})
}

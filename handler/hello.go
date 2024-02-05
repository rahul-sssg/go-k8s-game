package handler

import (
	"fmt"
	request "go-k8s-game/models/requests"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSum(ctx *gin.Context) {
	var request request.GetSumInput

	if err := ctx.ShouldBindQuery(&request); err != nil {
		msg := "invalid form body"
		data := gin.H{
			"message": msg,
		}
		ctx.JSON(http.StatusBadRequest, data)
		return
	}

	fmt.Println(request.A, " ", request.B)
	msg := fmt.Sprintf("User %s got sum %d", request.UserName, request.A+request.B)
	data := gin.H{
		"message": msg,
	}
	log.Printf(msg)
	ctx.JSON(http.StatusOK, data)
	return
}

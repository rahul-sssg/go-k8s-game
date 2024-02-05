package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

func GetHealth(ctx *gin.Context) {
	stats := make(map[string]interface{})
	stats["goRoutines"] = runtime.NumGoroutine()
	// stats["MySQL"] = base.DB.Stats()
	x, _ := json.Marshal(stats)

	fmt.Println("Health Stats:\t", string(x))
	ctx.JSON(http.StatusOK, stats)
}

package sd

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheck(c *gin.Context)  {
	    message := "ok"
	    c.String(http.StatusOK,"\n"+message)
}


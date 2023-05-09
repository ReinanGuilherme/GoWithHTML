package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {

	SetupRouterHome(r)

}

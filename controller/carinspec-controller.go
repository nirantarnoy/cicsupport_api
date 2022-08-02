package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.camelit.com/walofz/cicsupport-api/service"
)

type CarinspecController interface {
	NonconformList(ctx *gin.Context)
}

type carinspecController struct {
	carinspecService service.CarinspecSevice
	jwtService       service.JWTService
}

// NonconformList implements CarinspecController
func (db *carinspecController) NonconformList(ctx *gin.Context) {
	res := db.carinspecService.NonconformList()
	ctx.JSON(http.StatusOK, res)
}

func NewCarinspecController(carinspecService service.CarinspecSevice, jwtService service.JWTService) CarinspecController {
	return &carinspecController{carinspecService: carinspecService, jwtService: jwtService}
}

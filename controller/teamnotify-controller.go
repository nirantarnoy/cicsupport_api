package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gitlab.camelit.com/walofz/cicsupport-api/service"
)

type TeamNofiryController interface {
	FindEmpNotify(ctx *gin.Context)
}

type teamNotifyController struct {
	teamNotifyService service.TeamNotifyService
	jwtService        service.JWTService
}

// FindEmpNotify implements TeamNofiryController
func (db *teamNotifyController) FindEmpNotify(ctx *gin.Context) {
	empid, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := "No param"
		ctx.JSON(http.StatusBadRequest, res)
	}
	var teamnotify []entity.TeamNotify = db.teamNotifyService.FindEmpNotify(empid)
	ctx.JSON(http.StatusOK, teamnotify)
	// if (teamnotify == entity.TeamNotify{}) {
	// 	res := "Not found data"
	// 	ctx.JSON(http.StatusNotFound, res)
	// } else {
	// 	ctx.JSON(http.StatusOK, teamnotify)
	// }

}

func NewTeamNotifyController(teamnotyService service.TeamNotifyService, jwt service.JWTService) TeamNofiryController {
	return &teamNotifyController{teamNotifyService: teamnotyService, jwtService: jwt}
}

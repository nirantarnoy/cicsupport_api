package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gitlab.camelit.com/walofz/cicsupport-api/service"
)

type TeaminspectionitemController interface {
	FindInspectionItem(ctx *gin.Context)
}
type teaminspectionitemController struct {
	teaminspectionitemService service.TeaminspectionitemService
	jwtService                service.JWTService
}

// FindInspectionItem implements TeaminspectionitemController
func (db *teaminspectionitemController) FindInspectionItem(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := "No param"
		ctx.JSON(http.StatusBadRequest, res)
	}

	var teaminspection []entity.TeamInspectionItem = db.teaminspectionitemService.FindInspectionItem(id)
	if teaminspection == nil {
		res := "Data not found"
		ctx.JSON(http.StatusNotFound, res)
	} else {
		ctx.JSON(http.StatusOK, teaminspection)
	}

}

func NewIteminspectionitemController(teaminspectService service.TeaminspectionitemService, jwtService service.JWTService) TeaminspectionitemController {
	return &teaminspectionitemController{teaminspectionitemService: teaminspectService, jwtService: jwtService}
}

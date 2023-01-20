package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.camelit.com/walofz/cicsupport-api/dto"
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gitlab.camelit.com/walofz/cicsupport-api/helper"
	"gitlab.camelit.com/walofz/cicsupport-api/service"
)

type TeaminspectionitemController interface {
	FindInspectionItem(ctx *gin.Context)
	FindTransByEmp(ctx *gin.Context)
	FindTransHistoryByEmp(ctx *gin.Context)
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

func (db *teaminspectionitemController) FindTransByEmp(ctx *gin.Context) {
	var loginDTO dto.PersonTrans
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	s, err := strconv.ParseInt(loginDTO.TeamId, 0, 64)
	if err != nil {
		panic("format is invalid")
	}
	y, err := strconv.ParseInt(loginDTO.EmpId, 0, 64)
	if err != nil {
		panic("format is invalid")
	}
	res := db.teaminspectionitemService.FindTransByEmp(s, y)
	//res := "ok"
	ctx.JSON(http.StatusOK, res)

}

func (db *teaminspectionitemController) FindTransHistoryByEmp(ctx *gin.Context) {
	empid, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		res := "No param"
		ctx.JSON(http.StatusBadRequest, res)
	}
	// s, err := strconv.ParseInt(empid, 0, 64)
	// if err != nil {
	// 	panic("format is invalid")
	// }
	res := db.teaminspectionitemService.FindTransHistoryByEmp(empid)
	ctx.JSON(http.StatusOK, res)

}

func NewIteminspectionitemController(teaminspectService service.TeaminspectionitemService, jwtService service.JWTService) TeaminspectionitemController {
	return &teaminspectionitemController{teaminspectionitemService: teaminspectService, jwtService: jwtService}
}

package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.camelit.com/walofz/cicsupport-api/dto"
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gitlab.camelit.com/walofz/cicsupport-api/service"
)

type PlanController interface {
	FindPlanByTeam(ctx *gin.Context)
	AddInspection(ctx *gin.Context)
}

type planController struct {
	planService service.PlanService
	jwtService  service.JWTService
}

// AddInspection implements PlanController
func (db *planController) AddInspection(ctx *gin.Context) {
	var stock []dto.InspectionCreateDTO
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithError(400, err)
	}
	err = json.Unmarshal(body, &stock)
	if err != nil {
		ctx.AbortWithError(400, err)
	}
	// fmt.Printf("%v", stock)

	//	var inspectionDto []dto.InspectionCreateDTO
	fmt.Printf("%v", len(stock))
	errDto := ctx.Bind(&stock)
	if errDto != nil {
		fmt.Printf("%v", errDto)
		res := "Fail to precess request"
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		//	for i:=0
		res := db.planService.AddInspection(stock)
		ctx.JSON(http.StatusOK, res)
	}
}

// FindPlanByTeam implements PlanController
func (db *planController) FindPlanByTeam(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := "No param"
		ctx.JSON(http.StatusBadRequest, res)
	}

	var plans entity.Plan = db.planService.FindPlanByTeam(id)
	if (plans == entity.Plan{}) {
		res := "Data not found"
		ctx.JSON(http.StatusNotFound, res)
	} else {
		ctx.JSON(http.StatusOK, plans)
	}
}

func NewPlanController(planService service.PlanService, jwtService service.JWTService) PlanController {
	return &planController{planService: planService, jwtService: jwtService}
}

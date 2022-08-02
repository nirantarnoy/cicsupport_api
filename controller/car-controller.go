package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gitlab.camelit.com/walofz/cicsupport-api/dto"
	"gitlab.camelit.com/walofz/cicsupport-api/service"
)

type CarController interface {
	CreateCar(ctx *gin.Context)
	ListCarByEmpId(ctx *gin.Context)
}
type carController struct {
	carService service.CarService
	jwtService service.JWTService
}

// ListCarByEmpId implements CarController
func (db *carController) ListCarByEmpId(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := "Not param"
		ctx.JSON(http.StatusBadRequest, res)
	}
	res := db.carService.ListCarByEmpId(id)
	ctx.JSON(http.StatusOK, res)

}

// CreateCar implements CarController
func (db *carController) CreateCar(ctx *gin.Context) {
	//fmt.Print(ctx)
	var carDto dto.CarCreateDto
	errDto := ctx.ShouldBind(&carDto)
	if errDto != nil {
		fmt.Printf("%v", errDto)
		res := "Fail to process"
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := ctx.GetHeader("Authorization")
		userId := db.getUserIdByToken(authHeader)
		convertUserId, err := strconv.ParseUint(userId, 10, 64)

		if err == nil {
			carDto.CreatedBy = convertUserId
		}
		res := db.carService.CreateCar(carDto)
		ctx.JSON(http.StatusCreated, res)
	}

}

func NewCarController(carService service.CarService, jwtService service.JWTService) CarController {
	return &carController{carService: carService, jwtService: jwtService}
}
func (db *carController) getUserIdByToken(token string) string {
	aToken, err := db.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	} else {
		claims := aToken.Claims.(jwt.MapClaims)
		id := fmt.Sprintf("%v", claims["user_id"])
		return id
	}
}

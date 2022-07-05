package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gitlab.camelit.com/walofz/cicsupport-api/dto"
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gitlab.camelit.com/walofz/cicsupport-api/service"
)

type AreazoneController interface {
	Findall(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type areazoneController struct {
	areazoneService service.AreazoneService
	jwtService      service.JWTService
}

// Create implements AreazoneController
func (az *areazoneController) Create(ctx *gin.Context) {
	var areazoneCreateDto dto.AreazoneCreateDto
	errDto := ctx.ShouldBind(&areazoneCreateDto)
	if errDto != nil {
		fmt.Printf("%v", errDto)
		res := "Fail to process request naja"
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := ctx.GetHeader("Authorization")
		userId := az.getUserIdByToken(authHeader)
		convertUserId, err := strconv.ParseUint(userId, 10, 64)

		if err == nil {
			areazoneCreateDto.Created_by = convertUserId
		}
		res := az.areazoneService.CreateAreazone(areazoneCreateDto)
		ctx.JSON(http.StatusCreated, res)
	}
}

// Delete implements AreazoneController
func (az *areazoneController) Delete(ctx *gin.Context) {
	var areazone entity.Areazone
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := "No param id"
		ctx.JSON(http.StatusBadRequest, res)
	}

	areazone.ID = id
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := az.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	userId := fmt.Sprintf("%v", claims["user_id"])
	if az.areazoneService.CheckAllowed(userId, areazone.ID) {
		az.areazoneService.DeleteAreazone(areazone)
		res := "OK"
		ctx.JSON(http.StatusOK, res)
	} else {
		res := "Fail to get id"
		ctx.JSON(http.StatusBadRequest, res)
	}

}

// FindById implements AreazoneController
func (az *areazoneController) FindById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := "No param"
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var areazone entity.Areazone = az.areazoneService.FindById(id)
	if (areazone == entity.Areazone{}) {
		res := "Data not found"
		ctx.JSON(http.StatusNotFound, res)
	} else {
		ctx.JSON(http.StatusOK, areazone)
	}
}

// Findall implements AreazoneController
func (az *areazoneController) Findall(ctx *gin.Context) {
	var areazone []entity.Areazone = az.areazoneService.FindListAll()
	ctx.JSON(http.StatusOK, areazone)
}

// Update implements AreazoneController
func (az *areazoneController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := "No param id"
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var areazoneDto dto.AreazoneUpdateDto

	areazoneDto.ID = id

	errDTO := ctx.ShouldBind(&areazoneDto)
	if errDTO != nil {
		fmt.Print(errDTO.Error())
		res := "Fail to process"
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := az.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	//userId := fmt.Sprintf("%v", claims["user_id"])
	userId := "1"

	log.Print(claims)
	if az.areazoneService.CheckAllowed(userId, areazoneDto.ID) {
		az.areazoneService.UpdateAreazone(areazoneDto)
		res := "OK"
		ctx.JSON(http.StatusOK, res)
	} else {
		res := "Fail to get id"
		ctx.JSON(http.StatusBadRequest, res)
	}
}
func (ag *areazoneController) getUserIdByToken(token string) string {
	aToken, err := ag.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	} else {
		claims := aToken.Claims.(jwt.MapClaims)
		id := fmt.Sprintf("%v", claims["user_id"])
		return id
	}
}

func NewAreazoneController(areazoneService service.AreazoneService, jwtService service.JWTService) AreazoneController {
	return &areazoneController{areazoneService: areazoneService, jwtService: jwtService}
}

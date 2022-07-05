package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gitlab.camelit.com/walofz/cicsupport-api/dto"
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gitlab.camelit.com/walofz/cicsupport-api/service"
)

type AreagroupController interface {
	Listall(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type areagroupController struct {
	areagroupService service.AreagroupService
	jwtService       service.JWTService
}

// Create implements AreagroupController
func (ag *areagroupController) Create(ctx *gin.Context) {
	//fmt.Print(ctx)
	var areagroupCreateDto dto.AreagroupDTO
	errDto := ctx.ShouldBind(&areagroupCreateDto)
	if errDto != nil {
		fmt.Printf("%v", errDto)
		res := "Fail to process request naja"
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := ctx.GetHeader("Authorization")
		userId := ag.getUserIdByToken(authHeader)
		convertUserId, err := strconv.ParseUint(userId, 10, 64)

		if err == nil {
			areagroupCreateDto.Created_by = convertUserId
		}
		res := ag.areagroupService.Create(areagroupCreateDto)
		ctx.JSON(http.StatusCreated, res)
	}

}

// Delete implements AreagroupController
func (ag *areagroupController) Delete(ctx *gin.Context) {
	var aggroup entity.Areagroup
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := "No param id"
		ctx.JSON(http.StatusBadRequest, res)
	}

	aggroup.ID = id
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := ag.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	userId := fmt.Sprintf("%v", claims["user_id"])
	if ag.areagroupService.CheckAllowed(userId, aggroup.ID) {
		ag.areagroupService.Delete(aggroup)
		res := "OK"
		ctx.JSON(http.StatusOK, res)
	} else {
		res := "Fail to get id"
		ctx.JSON(http.StatusBadRequest, res)
	}

}

// Update implements AreagroupController
func (ag *areagroupController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := "No param id"
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var areagroupDto dto.AreagroupUpdateDTO

	areagroupDto.ID = id

	errDTO := ctx.ShouldBind(&areagroupDto)
	if errDTO != nil {
		fmt.Print(errDTO.Error())
		res := "Fail to process"
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := ag.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	userId := fmt.Sprintf("%v", claims["user_id"])

	if ag.areagroupService.CheckAllowed(userId, areagroupDto.ID) {
		ag.areagroupService.Update(areagroupDto)
		res := "OK"
		ctx.JSON(http.StatusOK, res)
	} else {
		res := "Fail to get id"
		ctx.JSON(http.StatusBadRequest, res)
	}
}

func NewAreagroupController(areagroupService service.AreagroupService, jwtService service.JWTService) AreagroupController {
	return &areagroupController{
		areagroupService: areagroupService,
		jwtService:       jwtService,
	}
}

// FindById implements AreagroupController
func (ag *areagroupController) FindById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := "No param"
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var areagroups entity.Areagroup = ag.areagroupService.FindById(id)
	if (areagroups == entity.Areagroup{}) {
		res := "Data not found"
		ctx.JSON(http.StatusNotFound, res)
	} else {
		ctx.JSON(http.StatusOK, areagroups)
	}
}

// Listall implements AreagroupController
func (areagroup_c *areagroupController) Listall(ctx *gin.Context) {
	var areagroups []entity.Areagroup = areagroup_c.areagroupService.ListAll()
	ctx.JSON(http.StatusOK, areagroups)
}

func (ag *areagroupController) getUserIdByToken(token string) string {
	aToken, err := ag.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	} else {
		claims := aToken.Claims.(jwt.MapClaims)
		id := fmt.Sprintf("%v", claims["user_id"])
		return id
	}
}

package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gitlab.camelit.com/walofz/cicsupport-api/helper"
	"gitlab.camelit.com/walofz/cicsupport-api/service"
)

type UserController interface {
	Profile(ctx *gin.Context)
	FindUserTeam(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

// FindUserTeam implements UserController
func (db *userController) FindUserTeam(ctx *gin.Context) {
	var userTeamId []entity.TeamMember
	userAd, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := "No param"
		ctx.JSON(http.StatusBadRequest, res)
	}

	userTeamId = db.userService.FindUserTeam(userAd)
	ctx.JSON(http.StatusOK, userTeamId)
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{userService: userService, jwtService: jwtService}
}

func (u *userController) Profile(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := u.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	user := u.userService.Profile(fmt.Sprintf("%v", claims["user_id"]))
	res := helper.BuildResponse(true, "OK", user)
	ctx.JSON(http.StatusOK, res)
}

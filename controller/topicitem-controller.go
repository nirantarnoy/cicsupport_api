package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gitlab.camelit.com/walofz/cicsupport-api/service"
)

type TopicItemController interface {
	FindTopicByPlan(ctx *gin.Context)
}

type topicItemController struct {
	topiceItemService service.TopicItemService
	jwtService        service.JWTService
}

// FindTopicByPlan implements TopicItemController
func (db *topicItemController) FindTopicByPlan(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := "No param"
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var topicitems []entity.TopicItem = db.topiceItemService.FindTopicByPlan(id)

	ctx.JSON(http.StatusOK, topicitems)

}

func NewTopicItemController(topicItemSevice service.TopicItemService, jwtService service.JWTService) TopicItemController {
	return &topicItemController{topiceItemService: topicItemSevice, jwtService: jwtService}
}

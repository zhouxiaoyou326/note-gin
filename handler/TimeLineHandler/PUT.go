package TimeLineHandler

import (
	"github.com/gin-gonic/gin"
	"note-gin/model"
	"note-gin/utils"
	"note-gin/view"
)

func Add(c *gin.Context) {
	time_line := model.TimeLine{}
	err := c.ShouldBind(&time_line)
	utils.ErrReport(err)

	time_line.Add()

	c.JSON(200, view.OkWithData("发表成功!", time_line))
}

func Update(c *gin.Context) {
	time_line := model.TimeLine{}
	err := c.ShouldBind(&time_line)
	utils.ErrReport(err)

	time_line.Update()

	c.JSON(200, view.OkWithData("修改成功!", time_line))
}

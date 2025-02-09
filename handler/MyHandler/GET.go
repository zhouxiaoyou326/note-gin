package MyHandler

import (
	"github.com/gin-gonic/gin"
	"note-gin/middleware/RedisClient"
	"note-gin/model"
	"note-gin/utils"
	"note-gin/view"
)

//Book
func GetAllBook(c *gin.Context) {

	books := RedisClient.GetAllBook()
	c.JSON(200, view.DataList{
		Items: books,
		Total: int64(len(books)),
	})
}

//Mood

func GetManyMood(c *gin.Context) {
	pageStr := c.Param("page")
	moods, total := model.MyMood{}.GetMany(utils.StrToInt(pageStr))
	c.JSON(200, view.DataList{
		Items: moods,
		Total: int64(total),
	})
}

func GetAccessRecord(c *gin.Context) {
	arr := RedisClient.GetAccessRecord()
	if len(arr) > 17 {
		arr = arr[:16]
	}
	c.JSON(200, view.DataList{
		Items: arr,
	})
}

package apis

import (
	"github.com/gin-gonic/gin"
	"sample/api/models"
)

func init(){

}

func Test01(c *gin.Context) {

	result := models.TestStruct{
		X1: "abc",
		X2: "123",
	}

	c.JSON(200, &result)
}


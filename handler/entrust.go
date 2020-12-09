package handler

import (
	"entrust/modles"
	"github.com/gin-gonic/gin"
)

// Entrust 下单接口
func Entrust(c *gin.Context) {

}

// Cancel 撤单接口
func Cancel(c *gin.Context) {

}

func OrderParamCheck(c *gin.Context) {
	entrust := modles.Entrust{}
	err := c.ShouldBindJSON(&entrust)
	if err != nil {
		c.AbortWithStatusJSON(500, "")
	}
}

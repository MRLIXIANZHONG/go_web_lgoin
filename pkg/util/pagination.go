package util

import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"my_web/pkg/setting"
)

func GetPage(c *gin.Context) int {
	result := 0//当前页码默认为0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}
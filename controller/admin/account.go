package admin

import (

	"github.com/gin-gonic/gin"
	"net/http"
	"web_test/model"
)

func (AdminController *admin) AccountIndex(c *gin.Context)  {
	//查询account信息
	err,account := model.AccountModel.AccountIndex()
	if err != "" {
		ReturnJSON(c, &Result{0, err, nil})
		return
	}

	c.HTML(http.StatusOK,"account/accounts.html",account)
}
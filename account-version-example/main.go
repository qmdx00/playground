package main

import (
	"account-version-example/biz"

	"github.com/gin-gonic/gin"
)

func main() {
	account := biz.NewAccountBiz()

	r := gin.Default()
	r.GET("/accounts/:uid", account.GetUserAccount)
	r.POST("/accounts/:uid/actions/init", account.InitUserAccount)
	r.POST("/accounts/:uid/actions/update", account.UpdateUserAccountAndCreateFlow)

	r.Run(":3000")
}

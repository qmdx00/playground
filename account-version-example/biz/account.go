package biz

import (
	"account-version-example/model"
	"account-version-example/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccountBiz struct {
	svc *service.AccountSvc
}

func NewAccountBiz() *AccountBiz {
	return &AccountBiz{svc: service.NewAccountSvc()}
}

// Example:
// curl --location --request GET 'http://127.0.0.1:3000/accounts/1'
func (a *AccountBiz) GetUserAccount(ctx *gin.Context) {
	uid := ctx.Param("uid")
	userId, _ := strconv.Atoi(uid)

	account, err := a.svc.GetUserAccount(int64(userId))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"status": "OK", "data": account})
}

// Example:
// curl --location --request POST 'http://127.0.0.1:3000/accounts/1/actions/init'
func (a *AccountBiz) InitUserAccount(ctx *gin.Context) {
	uid := ctx.Param("uid")
	userId, err := strconv.Atoi(uid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}

	account, err := a.svc.InitUserAccount(int64(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "OK", "data": account})
}

type UpdateBody struct {
	Amount   float64 `json:"amount"`
	FlowType int32   `json:"type"`
	BizNo    string  `json:"bizNo"`
}

// Example:
// curl --location --request POST 'http://127.0.0.1:3000/accounts/1/actions/update' -d '{"amount":1.23,"type":1,"bizNo":"xxxxxxxx"}'
func (a *AccountBiz) UpdateUserAccountAndCreateFlow(ctx *gin.Context) {
	uid := ctx.Param("uid")
	userId, err := strconv.Atoi(uid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}

	params := new(UpdateBody)
	if err := ctx.ShouldBindBodyWithJSON(params); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}

	account, accountFlow, err := a.svc.UpdateUserAccountAndCreateFlow(int64(userId), params.Amount, params.BizNo, model.AccountFlowType(params.FlowType))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data": gin.H{
			"account": account,
			"flow":    accountFlow,
		},
	})
}

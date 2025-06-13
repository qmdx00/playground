package model

import (
	"time"
)

type (
	AccountStatus int32

	AccountFlowType int32
)

const (
	AccountNormal AccountStatus = iota + 1
	AccountFreeze
)

const (
	AccountFlowRecharge AccountFlowType = iota + 1
	AccountFlowConsume
	AccountFlowRefund
	AccountFlowWithdraw
)

type Account struct {
	ID        int64     `db:"id" json:"id"`
	UserID    int64     `db:"user_id" json:"userId"`  // 用户id
	Balance   float64   `db:"balance" json:"balance"` // 当前余额，精确到分
	Version   int       `db:"version" json:"version"` // 版本号
	Status    int32     `db:"status" json:"status"`   // 账户状态（1：正常、2：冻结）
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}

type AccountFlow struct {
	ID            int64     `db:"id" json:"id"`
	FlowNo        string    `db:"flow_no" json:"flowNo"`               // 流水号
	AccountID     int64     `db:"account_id" json:"accountId"`         // 关联的账户id
	Amount        float64   `db:"amount" json:"amount"`                // 变动金额（正：进账，负：出账）
	BalanceBefore float64   `db:"balance_before" json:"balanceBefore"` // 变动前余额
	BalanceAfter  float64   `db:"balance_after" json:"balanceAfter"`   // 变动后余额
	Type          int32     `db:"type" json:"type"`                    // 流水类型（1：充值、2：消费、3：退款、4：提现）
	BizNo         string    `db:"biz_no" json:"bizNo"`                 // 业务单号
	VersionSeq    int       `db:"version_seq" json:"versionSeq"`       // 关联账户的版本号（用于追溯）
	CreatedAt     time.Time `db:"created_at" json:"createdAt"`
}

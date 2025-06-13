package service

import (
	"account-version-example/data"
	"account-version-example/model"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AccountSvc struct {
	ds *sqlx.DB
}

func NewAccountSvc() *AccountSvc {
	return &AccountSvc{ds: data.InitMySQL()}
}

func (s *AccountSvc) GetUserAccount(userId int64) (*model.Account, error) {
	account := new(model.Account)
	err := s.ds.Get(account, "SELECT * FROM account WHERE user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (s *AccountSvc) InitUserAccount(userId int64) (*model.Account, error) {
	account := &model.Account{
		UserID:    userId,
		Balance:   0,
		Version:   0,
		Status:    int32(model.AccountNormal),
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
	}

	result, err := s.ds.NamedExec(`INSERT INTO account (user_id, balance, version, status, created_at, updated_at) VALUES (:user_id, :balance, :version, :status, :created_at, :updated_at)`, account)
	if err != nil {
		return nil, err
	}

	account.ID, _ = result.LastInsertId()
	return account, nil
}

func (s *AccountSvc) UpdateUserAccountAndCreateFlow(userId int64, amount float64, bizNo string, flowType model.AccountFlowType) (*model.Account, *model.AccountFlow, error) {
	// NOTE: 在主从架构下，需在主库执行事务，避免主从延迟导致的版本号不一致，导致大量失败的事务回滚
	tx, err := s.ds.Beginx()
	if err != nil {
		return nil, nil, err
	}
	defer tx.Commit()

	// NOTE: 在高频写入场景下，需要加 FOR UPDATE，否则会导致大量失败的事务回滚
	// NOTE: 在读多写少且低频写入的场景下，可以不加 FOR UPDATE，避免频繁的锁竞争（X 型行锁）
	account := new(model.Account)
	err = tx.Get(account, "SELECT * FROM account WHERE user_id = ? FOR UPDATE", userId) // NOTE: 可只查询必要字段，使用联合索引来避免回表操作
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	result, err := tx.NamedExec(`UPDATE account SET balance = balance + :amount, version = version + 1 WHERE id = :id AND version = :version`, map[string]any{"amount": amount, "id": account.ID, "version": account.Version})
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	// NOTE: 抛出错误，由客户端主动重试，此处也可以使用自动重试，但需要考虑重试次数和重试间隔
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		tx.Rollback()
		return nil, nil, errors.New("version conflict, please retry")
	}

	accountFlow := &model.AccountFlow{
		FlowNo:        uuid.NewString(),
		AccountID:     account.ID,
		Amount:        amount,
		BalanceBefore: account.Balance,
		BalanceAfter:  account.Balance + amount,
		Type:          int32(flowType),
		BizNo:         bizNo,
		VersionSeq:    account.Version + 1,
		CreatedAt:     time.Now().Local(),
	}

	result, err = tx.NamedExec(`INSERT INTO account_flow (flow_no, account_id, amount, balance_before, balance_after, type, biz_no, version_seq, created_at) VALUES (:flow_no, :account_id, :amount, :balance_before, :balance_after, :type, :biz_no, :version_seq, :created_at)`, accountFlow)
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	accountFlow.ID, _ = result.LastInsertId()
	return account, accountFlow, nil
}

package dbmodel

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
)

type CheckUserAlive struct {
	ID                     int    `json:"id" gorm:"column:id"`
	Username               string `json:"username" gorm:"column:username"`
	Password               string `json:"password" gorm:"column:password"`
	DayLoginFirstTime      int    `json:"day_login_first_time" gorm:"column:day_login_first_time"`
	SendID                 int    `json:"send_id" gorm:"column:send_id"`
	SendEmailActionTimeout int    `json:"send_email_action_timeout" gorm:"column:send_email_action_timeout"`
	Salt                   string `json:"salt" gorm:"column:salt"`
	SuperSendConnInfoId    int    `json:"super_send_conn_info_id" gorm:"column:super_send_conn_info_id"`
}
type CheckUserAliveJoinConnInfo struct {
	ID                       int    `json:"id" gorm:"column:id"`
	Username                 string `json:"username" gorm:"column:username"`
	Password                 string `json:"password" gorm:"column:password"`
	DayLoginFirstTime        int    `json:"day_login_first_time" gorm:"column:day_login_first_time"`
	SendID                   int    `json:"send_id" gorm:"column:send_id"`
	SendEmailActionTimeout   int    `json:"send_email_action_timeout" gorm:"column:send_email_action_timeout"`
	Salt                     string `json:"salt" gorm:"column:salt"`
	SuperSendConnInfoId      int    `json:"super_send_conn_info_id" gorm:"column:super_send_conn_info_id"`
	SuperSendConnInfoAddress string `json:"super_send_conn_info_address" gorm:"column:super_send_conn_info_address"`
}

func (m *CheckUserAlive) TableName() string {
	return "check_user_alive"
}

type CheckUserAliveDao struct {
	sourceDB  *gorm.DB
	replicaDB []*gorm.DB
	m         *CheckUserAlive
}

func NewCheckUserAliveDao(ctx context.Context, dbs ...*gorm.DB) *CheckUserAliveDao {
	dao := new(CheckUserAliveDao)
	switch len(dbs) {
	case 0:
		panic("database connection required")
	case 1:
		dao.sourceDB = dbs[0]
		dao.replicaDB = []*gorm.DB{dbs[0]}
	default:
		dao.sourceDB = dbs[0]
		dao.replicaDB = dbs[1:]
	}
	dao.m = new(CheckUserAlive)
	return dao
}

func (d *CheckUserAliveDao) Create(ctx context.Context, obj *CheckUserAlive) error {
	err := d.sourceDB.Model(d.m).Create(&obj).Error
	if err != nil {
		return fmt.Errorf("CheckUserAliveDao: %w", err)
	}
	return nil
}

func (d *CheckUserAliveDao) Get(ctx context.Context, fields, where string, args ...interface{}) (*CheckUserAlive, error) {
	items, err := d.List(ctx, fields, where, 0, 1, args...)
	if err != nil {
		return nil, fmt.Errorf("CheckUserAliveDao: Get where=%s: %w", where, err)
	}
	if len(items) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &items[0], nil
}

func (d *CheckUserAliveDao) List(ctx context.Context, fields, where string, offset, limit interface{}, args ...interface{}) ([]CheckUserAlive, error) {
	var results []CheckUserAlive
	err := d.replicaDB[rand.Intn(len(d.replicaDB))].Model(d.m).
		Select(fields).Where(where, args...).Offset(offset.(int)).Limit(limit.(int)).Find(&results).Error
	if err != nil {
		return nil, fmt.Errorf("CheckUserAliveDao: List where=%s: %w", where, err)
	}
	return results, nil
}
func (d *CheckUserAliveDao) ListLeftJoinSuperSendInfo(ctx context.Context, fields, where string, offset, limit interface{}, args ...interface{}) ([]CheckUserAliveJoinConnInfo, error) {
	var results []CheckUserAliveJoinConnInfo
	err := d.replicaDB[rand.Intn(len(d.replicaDB))].Model(d.m).Joins("LEFT JOIN super_send_conn_info ON super_send_conn_info.id = check_user_alive.super_send_conn_info_id").
		Select(fields).Where(where, args...).Offset(offset.(int)).Limit(limit.(int)).Find(&results).Error
	if err != nil {
		return nil, fmt.Errorf("CheckUserAliveDao: List where=%s: %w", where, err)
	}
	return results, nil
}
func (d *CheckUserAliveDao) ListAll(ctx context.Context, fields, where string, args ...interface{}) ([]CheckUserAlive, error) {
	var results []CheckUserAlive
	err := d.replicaDB[rand.Intn(len(d.replicaDB))].Model(d.m).
		Select(fields).Where(where, args...).Find(&results).Error
	if err != nil {
		return nil, fmt.Errorf("CheckUserAliveDao: List where=%s: %w", where, err)
	}
	return results, nil
}
func (d *CheckUserAliveDao) Update(ctx context.Context, where string, update map[string]interface{}, args ...interface{}) error {
	err := d.sourceDB.Model(d.m).Where(where, args...).
		Updates(update).Error
	if err != nil {
		return fmt.Errorf("CheckUserAliveDao:Update where=%s: %w", where, err)
	}
	return nil
}

func (d *CheckUserAliveDao) Delete(ctx context.Context, where string, args ...interface{}) error {
	if len(where) == 0 {
		return gorm.ErrInvalidData
	}
	if err := d.sourceDB.Where(where, args...).Delete(d.m).Error; err != nil {
		return fmt.Errorf("CheckUserAliveDao: Delete where=%s: %w", where, err)
	}
	return nil
}
func (d *CheckUserAliveDao) Count(ctx context.Context, fields, where string, args ...interface{}) (count int64, err error) {
	err = d.replicaDB[rand.Intn(len(d.replicaDB))].Model(d.m).
		Select(fields).Where(where, args...).Count(&count).Error
	if err != nil {
		return 0, fmt.Errorf("CheckUserAliveDao: List where=%s: %w", where, err)
	}
	return
}

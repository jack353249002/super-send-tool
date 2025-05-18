package dbmodel

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
)

type EtcdBridgeConnInfo struct {
	ID       int    `json:"id" gorm:"column:id"`
	Address  string `json:"address" gorm:"column:address"`
	IsSsl    int    `json:"is_ssl" gorm:"column:is_ssl"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
	Online   int    `json:"online" gorm:"column:online"`
}

func (m *EtcdBridgeConnInfo) TableName() string {
	return "etcd_bridge_conn_info"
}

type EtcdBridgeConnInfoDao struct {
	sourceDB  *gorm.DB
	replicaDB []*gorm.DB
	m         *EtcdBridgeConnInfo
}

func NewEtcdBridgeConnInfoDao(ctx context.Context, dbs ...*gorm.DB) *EtcdBridgeConnInfoDao {
	dao := new(EtcdBridgeConnInfoDao)
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
	dao.m = new(EtcdBridgeConnInfo)
	return dao
}

func (d *EtcdBridgeConnInfoDao) Create(ctx context.Context, obj *EtcdBridgeConnInfo) error {
	err := d.sourceDB.Model(d.m).Create(&obj).Error
	if err != nil {
		return fmt.Errorf("EtcdBridgeConnInfoDao: %w", err)
	}
	return nil
}

func (d *EtcdBridgeConnInfoDao) Get(ctx context.Context, fields, where string, args ...interface{}) (*EtcdBridgeConnInfo, error) {
	items, err := d.List(ctx, fields, where, 0, 1, args...)
	if err != nil {
		return nil, fmt.Errorf("EtcdBridgeConnInfoDao: Get where=%s: %w", where, err)
	}
	if len(items) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &items[0], nil
}

func (d *EtcdBridgeConnInfoDao) List(ctx context.Context, fields, where string, offset, limit interface{}, args ...interface{}) ([]EtcdBridgeConnInfo, error) {
	var results []EtcdBridgeConnInfo
	err := d.replicaDB[rand.Intn(len(d.replicaDB))].Model(d.m).
		Select(fields).Where(where, args...).Offset(offset.(int)).Limit(limit.(int)).Find(&results).Error
	if err != nil {
		return nil, fmt.Errorf("EtcdBridgeConnInfoDao: List where=%s: %w", where, err)
	}
	return results, nil
}

func (d *EtcdBridgeConnInfoDao) Update(ctx context.Context, where string, update map[string]interface{}, args ...interface{}) error {
	err := d.sourceDB.Model(d.m).Where(where, args...).
		Updates(update).Error
	if err != nil {
		return fmt.Errorf("EtcdBridgeConnInfoDao:Update where=%s: %w", where, err)
	}
	return nil
}
func (d *EtcdBridgeConnInfoDao) Count(ctx context.Context, fields, where string, args ...interface{}) (count int64, err error) {
	err = d.replicaDB[rand.Intn(len(d.replicaDB))].Model(d.m).
		Select(fields).Where(where, args...).Count(&count).Error
	if err != nil {
		return 0, fmt.Errorf("EtcdBridgeConnInfoDao: List where=%s: %w", where, err)
	}
	return
}
func (d *EtcdBridgeConnInfoDao) Delete(ctx context.Context, where string, args ...interface{}) error {
	if len(where) == 0 {
		return gorm.ErrInvalidData
	}
	if err := d.sourceDB.Where(where, args...).Delete(d.m).Error; err != nil {
		return fmt.Errorf("EtcdBridgeConnInfoDao: Delete where=%s: %w", where, err)
	}
	return nil
}

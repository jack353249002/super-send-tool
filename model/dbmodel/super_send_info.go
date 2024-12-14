package dbmodel

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
)

type SuperSendConnInfo struct {
	ID       int    `json:"id" gorm:"column:id"`
	Address  string `json:"address" gorm:"column:address"`
	IsSsl    int    `json:"is_ssl" gorm:"column:is_ssl"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
	Online   uint8  `json:"online"`
	Token    string `json:"token"`
}

func (m *SuperSendConnInfo) TableName() string {
	return "super_send_conn_info"
}

type SuperSendConnInfoDao struct {
	sourceDB  *gorm.DB
	replicaDB []*gorm.DB
	m         *SuperSendConnInfo
}

func NewSuperSendConnInfoDao(ctx context.Context, dbs ...*gorm.DB) *SuperSendConnInfoDao {
	dao := new(SuperSendConnInfoDao)
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
	dao.m = new(SuperSendConnInfo)
	return dao
}

func (d *SuperSendConnInfoDao) Create(ctx context.Context, obj *SuperSendConnInfo) error {
	err := d.sourceDB.Model(d.m).Create(&obj).Error
	if err != nil {
		return fmt.Errorf("SuperSendConnInfoDao: %w", err)
	}
	return nil
}

func (d *SuperSendConnInfoDao) Get(ctx context.Context, fields, where string, args ...interface{}) (*SuperSendConnInfo, error) {
	items, err := d.List(ctx, fields, where, 0, 1, args...)
	if err != nil {
		return nil, fmt.Errorf("SuperSendConnInfoDao: Get where=%s: %w", where, err)
	}
	if len(items) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &items[0], nil
}

func (d *SuperSendConnInfoDao) List(ctx context.Context, fields, where string, offset, limit int, args ...interface{}) ([]SuperSendConnInfo, error) {
	var results []SuperSendConnInfo
	err := d.replicaDB[rand.Intn(len(d.replicaDB))].Model(d.m).
		Select(fields).Where(where, args...).Order("id desc").Offset(offset).Limit(limit).Find(&results).Error
	if err != nil {
		return nil, fmt.Errorf("SuperSendConnInfoDao: List where=%s: %w", where, err)
	}
	return results, nil
}
func (d *SuperSendConnInfoDao) ListAll(ctx context.Context, fields, where string, args ...interface{}) ([]SuperSendConnInfo, error) {
	var results []SuperSendConnInfo
	err := d.replicaDB[rand.Intn(len(d.replicaDB))].Model(d.m).
		Select(fields).Where(where, args...).Order("id desc").Find(&results).Error
	if err != nil {
		return nil, fmt.Errorf("SuperSendConnInfoDao: List where=%s: %w", where, err)
	}
	return results, nil
}
func (d *SuperSendConnInfoDao) Count(ctx context.Context, fields, where string, args ...interface{}) (count int64, err error) {
	err = d.replicaDB[rand.Intn(len(d.replicaDB))].Model(d.m).
		Select(fields).Where(where, args...).Count(&count).Error
	if err != nil {
		return 0, fmt.Errorf("SuperSendConnInfoDao: List where=%s: %w", where, err)
	}
	return
}

func (d *SuperSendConnInfoDao) Update(ctx context.Context, where string, update map[string]interface{}, args ...interface{}) error {
	err := d.sourceDB.Model(d.m).Where(where, args...).
		Updates(update).Error
	if err != nil {
		return fmt.Errorf("SuperSendConnInfoDao:Update where=%s: %w", where, err)
	}
	return nil
}

func (d *SuperSendConnInfoDao) Delete(ctx context.Context, where string, args ...interface{}) error {
	if len(where) == 0 {
		return gorm.ErrMissingWhereClause
	}
	if err := d.sourceDB.Where(where, args...).Delete(d.m).Error; err != nil {
		return fmt.Errorf("SuperSendConnInfoDao: Delete where=%s: %w", where, err)
	}
	return nil
}

package install

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"super-send-tool/config/baseconfig"
)

var InstallError = errors.New("install error")
var LockInfo Lock

type Lock struct {
	DataBase bool
	Root     bool
}

func init() {
	LockInfo.DataBase = false
	LockInfo.Root = false
	data, err := os.ReadFile("lock")
	if err == nil {
		err = json.Unmarshal(data, &LockInfo)
	}
}

// executeSQLFile 执行 SQL 文件中的 SQL 语句
func executeSQLFile() error {
	//db, err := sql.Open("sqlite3", baseconfig.CONFIG.DBPath)
	sourceName := fmt.Sprintf("file:%s?_foreign_keys=on", baseconfig.CONFIG.DBPath)
	db, err := sql.Open("sqlite3", sourceName)
	if err != nil {
		return err
	}
	defer func() {
		if err := db.Close(); err != nil {
			return
		}
	}()
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations", // SQL 迁移文件路径
		"sqlite3", driver,   // SQLite URL 格式
	)
	if err != nil {
		return fmt.Errorf("创建迁移实例失败: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("执行迁移失败: %v", err)
	}

	fmt.Println("数据库迁移成功")
	return nil
}
func InstallDataBase() error {
	_, err := os.Stat("lock")
	//安装数据库
	executeSQLFile()
	if os.IsNotExist(err) {
		LockInfo.DataBase = true
		LockInfo.Root = false
		data, err := json.Marshal(LockInfo)
		if err == nil {
			err = os.WriteFile("lock", data, os.ModePerm)
		}
		return err
	} else {
		return InstallError
	}
}
func SetLockInfo(key string, val bool) {
	_, err := os.Stat("lock")
	if os.IsNotExist(err) {
		if key == "root" {
			LockInfo.Root = val
		}
		if key == "database" {
			LockInfo.DataBase = val
		}
		data, err := json.Marshal(LockInfo)
		if err == nil {
			err = os.WriteFile("lock", data, os.ModePerm)
		}
	} else {
		data, err := os.ReadFile("lock")
		if err == nil {
			err = json.Unmarshal(data, &LockInfo)
			if err == nil {
				if key == "root" {
					LockInfo.Root = val
				}
				if key == "database" {
					LockInfo.DataBase = val
				}
				data, err := json.Marshal(LockInfo)
				if err == nil {
					err = os.WriteFile("lock", data, os.ModePerm)
				}
			}
		}
	}
}
func CheckRootLock() bool {
	return LockInfo.Root
}
func CheckDataBaseLock() bool {
	return LockInfo.DataBase
}

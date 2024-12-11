package install

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"super-send-tool/db"
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
func executeSQLFile(db *gorm.DB, filePath string) error {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	db.Exec(string(content))
	return err
}
func InstallDataBase() error {
	_, err := os.Stat("lock")
	if os.IsNotExist(err) {
		//安装数据库
		executeSQLFile(db.Db, "sql/script.sql")
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

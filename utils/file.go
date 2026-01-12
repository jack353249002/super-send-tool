package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
)

// CreateFolderByMD5 创建一个以字符串的MD5哈希值命名的文件夹
func CreateFolderByMD5(savePath, input string) (string, string, error) {
	hash := md5.New()
	hash.Write([]byte(input))
	hashed := hex.EncodeToString(hash.Sum(nil))
	folderPath := savePath + "/" + hashed

	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err := os.MkdirAll(folderPath, 0755)
		if err != nil {
			return "", "", fmt.Errorf("failed to create folder: %w", err)
		}
	}

	return folderPath, hashed, nil
}

// md5加密
func MD5(info, salt string) string {
	data := []byte(info + salt)
	s := fmt.Sprintf("%x", md5.Sum(data))
	return s
}

// FolderExistsByMD5 检查以字符串的MD5哈希值命名的文件夹是否存在
func FolderExistsByMD5(savePath, input string) (bool, error) {
	hash := md5.New()
	hash.Write([]byte(input))
	hashed := hex.EncodeToString(hash.Sum(nil))
	folderPath := filepath.Join(savePath, hashed)

	_, err := os.Stat(folderPath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// WriteBytesToFile 将字节数据写入到指定路径的文件中
func WriteBytesToFile(filePath string, data []byte) error {
	err := ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write bytes to file: %w", err)
	}
	return nil
}

// GenerateSalt 生成指定长度的盐值
func GenerateSalt(length int) (string, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

package hashpwd

import (
	"golang.org/x/crypto/bcrypt"
)

// 哈希密码
func HashPassword(password string) (string, error) {
	// 生成哈希值，第二个参数是cost因子，值越大越安全但也越耗时
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// 验证密码
func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

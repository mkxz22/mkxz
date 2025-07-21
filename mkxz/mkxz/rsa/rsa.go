package rsa

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

// AES对称加密示例
func aesEncrypt(plaintext []byte, key []byte) ([]byte, error) {
	// 创建AES块
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 使用CBC模式，需要一个初始化向量
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// 加密
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

// AES对称解密示例
func aesDecrypt(ciphertext []byte, key []byte) ([]byte, error) {
	// 创建AES块
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 检查密文长度
	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("密文长度过短")
	}

	// 提取初始化向量
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// 解密
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return ciphertext, nil
}

// RSA非对称加密示例
func rsaEncrypt(plaintext []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	// 使用OAEP填充进行加密
	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, plaintext, nil)
	if err != nil {
		return nil, err
	}

	return ciphertext, nil
}

// RSA非对称解密示例
func rsaDecrypt(ciphertext []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	// 使用OAEP填充进行解密
	hash := sha256.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func main() {
	// AES示例
	plaintext := []byte("Hello, AES!")
	// AES-256需要32字节密钥
	key := []byte("this-is-a-secret-key-32-bytes-long-1234")

	ciphertext, err := aesEncrypt(plaintext, key)
	if err != nil {
		fmt.Println("AES加密错误:", err)
		return
	}
	fmt.Println("AES密文:", base64.StdEncoding.EncodeToString(ciphertext))

	decrypted, err := aesDecrypt(ciphertext, key)
	if err != nil {
		fmt.Println("AES解密错误:", err)
		return
	}
	fmt.Println("AES明文:", string(bytes.TrimRight(decrypted, "\x00")))

	// RSA示例
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("生成RSA密钥对错误:", err)
		return
	}
	publicKey := &privateKey.PublicKey

	message := []byte("Hello, RSA!")
	encrypted, err := rsaEncrypt(message, publicKey)
	if err != nil {
		fmt.Println("RSA加密错误:", err)
		return
	}
	fmt.Println("RSA密文:", base64.StdEncoding.EncodeToString(encrypted))

	decoded, err := rsaDecrypt(encrypted, privateKey)
	if err != nil {
		fmt.Println("RSA解密错误:", err)
		return
	}
	fmt.Println("RSA明文:", string(decoded))
}

package oss

import (
	"context"
	"flag"
	"log"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
)

func Oss(region, dst, bucketName, filename, accessKeyid, accesskeysecret string) {
	// 解析命令行参数
	flag.Parse()

	// 加载默认配置并设置凭证提供者和区域
	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyid, accesskeysecret)).
		WithRegion(region)

	// 创建OSS客户端
	client := oss.NewClient(cfg)

	// 创建上传管理器
	u := client.NewUploader()

	// 定义本地文件路径，需要替换为您的实际本地文件路径和文件名称
	localFile := dst

	// 执行上传文件的操作
	result, err := u.UploadFile(context.TODO(),
		&oss.PutObjectRequest{
			Bucket: oss.Ptr(bucketName),
			Key:    oss.Ptr(filename)},
		localFile)
	if err != nil {
		log.Fatalf("failed to upload file %v", err)
	}

	// 打印上传文件的结果
	log.Printf("upload file result:%#v\n", result)
}

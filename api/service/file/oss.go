//Package file
//@ Author: Gao YuanMing
//@ Data: 2019/12/17 9:09 下午
//@ Description:

package file

import (
	"api/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
)

var (
	client *oss.Client
	bucket *oss.Bucket
	err    error
)

func init() {
	// 创建OSSClient实例
	client, err = oss.New(config.OSSEndpoint, config.OSSAccessKeyID, config.OSSAccessKeySecret)
	if err != nil {
		log.Println("OSS client create error:", err)
		return
	}
	// 创建存储空间
	bucket, err = client.Bucket(config.OSSBucketName)
	if err != nil {
		log.Println("OSS Bucket create error:", err)
		return
	}
}

func GetBucket() *oss.Bucket {
	return bucket
}

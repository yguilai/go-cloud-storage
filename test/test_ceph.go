package main

import (
	"fmt"
	"github.com/yguilai/go-cloud-storage/store/ceph"

	"gopkg.in/amz.v1/s3"
)

func __main() {
	bucket := ceph.GetCephBucket("testbucket1")

	// // 创建一个新的bucket
	err := bucket.PutBucket(s3.PublicRead)
	fmt.Printf("create bucket err: %v\n", err)

	// d, _ := bucket.Get("/ceph/866cc7c87c9b612dd8904d2c5dd07d6f6c22b834")
	// tmpFile, _ := os.Create("/tmp/test_file")
	// tmpFile.Write(d)

	// 查询这个bucket下面指定条件的object keys
	res, _ := bucket.List("", "", "", 99)
	fmt.Printf("%+v\n", res)

	// 新上传一个对象
	err = bucket.Put("./tmp/Mysql分库分表.md", []byte("just for test"), "octet-stream", s3.PublicRead)
	fmt.Printf("upload err: %+v\n", err)

	// 查询这个bucket下面指定条件的object keys
	res, err = bucket.List("", "", "", 100)
	fmt.Printf("object keys: %+v\n", res)
}

package config

import (
	cmn "github.com/yguilai/go-cloud-storage/common"
)

const (
	// TempLocalRootDir : 本地临时存储地址的路径
	TempLocalRootDir = "/data/fileserver/"
	// TempPartRootDir : 分块文件在本地临时存储地址的路径
	TempPartRootDir = "/data/fileserver_part/"
	// CephRootDir : Ceph的存储路径prefix
	CephRootDir = "/ceph"
	// KODORootDir : KODO的存储路径prefix
	KODORootDir = "kodo/"
	// CurrentStoreType : 设置当前文件的存储类型
	CurrentStoreType = cmn.StoreLocal
)

package rpc

import (
	"context"
	cfg "github.com/yguilai/go-cloud-storage/service/upload/config"
	upProto "github.com/yguilai/go-cloud-storage/service/upload/proto"
)

// Upload : upload结构体
type Upload struct{}

// UploadEntry : 获取上传入口
func (u *Upload) UploadEntry(
	ctx context.Context,
	req *upProto.ReqEntry,
	res *upProto.RespEntry) error {

	res.Entry = cfg.UploadEntry
	return nil
}

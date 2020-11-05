package kodo

import (
	"context"
	"github.com/qiniu/api.v7/v7/auth"
	"github.com/qiniu/api.v7/v7/storage"
	"io"
	"time"
)

const (
	// TODO 上传前修改a/s key
	ACCESS_KEY  = ""
	SECRET_KEY  = ""
	BUCKET      = "your bucket"
	DOMAIN = "your domain"
)

var (
	uploader *storage.FormUploader
)

func init() {
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuanan
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false
	uploader = storage.NewFormUploader(&cfg)
}

func UploadToQiniu(key string, data io.Reader, size int64) (*storage.PutRet, error) {
	pp := storage.PutPolicy{
		Scope: BUCKET,
	}

	mac := auth.New(ACCESS_KEY, SECRET_KEY)
	upToken := pp.UploadToken(mac)

	ret := &storage.PutRet{}
	err := uploader.Put(context.Background(), ret, upToken, key, data, size, nil)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func DownloadUrl(key string) string {
	return storage.MakePrivateURL(auth.New(ACCESS_KEY, SECRET_KEY), DOMAIN, key, time.Now().Add(3*time.Minute).Unix())
}
package process

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github.com/yguilai/go-cloud-storage/mq"
	dbcli "github.com/yguilai/go-cloud-storage/service/dbproxy/client"
	"github.com/yguilai/go-cloud-storage/store/kodo"
)

// Transfer : 处理文件转移
func Transfer(msg []byte) bool {
	log.Println(string(msg))

	pubData := mq.TransferData{}
	err := json.Unmarshal(msg, &pubData)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	fin, err := os.Open(pubData.CurLocation)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	fi, err := fin.Stat()
	if err != nil {
		log.Println(err.Error())
		return false
	}

	ret, err := kodo.UploadToQiniu(pubData.DestLocation, bufio.NewReader(fin), fi.Size())

	if err != nil {
		log.Println(err.Error())
		return false
	}

	resp, err := dbcli.UpdateFileLocation(ret.Hash, ret.Key)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	if !resp.Suc {
		log.Println("更新数据库异常，请检查:" + ret.Hash)
		return false
	}
	return true
}

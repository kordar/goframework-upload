package goframework_upload

import (
	"github.com/kordar/godb"
	"github.com/kordar/goupload"
)

var (
	uploadpool = godb.NewDbPool()
)

func GetUploaderClient(db string) goupload.BucketUploader {
	return uploadpool.Handle(db).(goupload.BucketUploader)
}

// AddUploaderInstance 添加uploader
func AddUploaderInstance(db string, client goupload.BucketUploader) error {
	ins := NewUploaderIns(db, client)
	return uploadpool.Add(ins)
}

// RemoveUploaderInstance 移除uploader
func RemoveUploaderInstance(db string) {
	uploadpool.Remove(db)
}

// HasUploaderInstance uploader句柄是否存在
func HasUploaderInstance(db string) bool {
	return uploadpool != nil && uploadpool.Has(db)
}

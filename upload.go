package goframework_upload

import "github.com/kordar/goupload"

type UploaderIns struct {
	name string
	ins  goupload.BucketUploader
}

func NewUploaderIns(name string, client goupload.BucketUploader) *UploaderIns {
	return &UploaderIns{name: name, ins: client}
}

func (c UploaderIns) GetName() string {
	return c.name
}

func (c UploaderIns) GetInstance() interface{} {
	return c.ins
}

func (c UploaderIns) Close() error {
	return nil
}

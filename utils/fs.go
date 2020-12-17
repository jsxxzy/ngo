package utils

import (
	"github.com/cavaliercoder/grab"
)

// DownloadFile 下载文件
func DownloadFile(remote, localpath string) error {

	_, err := grab.Get(localpath, remote)
	if err != nil {
		return err
	}
	return nil
}

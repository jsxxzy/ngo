package utils

import (
	"io"
	"net/http"
	"os"
)

// DownloadFile 下载文件
func DownloadFile(remote, localpath string) error {

	// Get the data
	resp, err := http.Get(remote)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(localpath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err

}

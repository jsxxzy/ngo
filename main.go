package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/jsxxzy/ngo/utils"
)

//go:generate go-bindata -o data.go vscode.bat

var (
	goMirrorURL    = "https://studygolang.com"                // `golang` 镜像站
	gitDownloadURL string                                     // `git` 下载地址
	vscodeBatFile  string                      = "vscode.bat" // `vscode` 批处理安装文件
)

func main() {
	setup()
}

func setup() {
	installGO()
	installGIT()
	installVSCODE()
}

func installGO() {

	// ######## 创建环境变量
	utils.SetGlobalENVBool("GO111MODULE", "on")
	utils.SetGlobalENVBool("GOPROXY", "https://goproxy.io,direct")
	// ########

	var downloadURL, filename = getGolangDownloadURL()

	if !utils.Check(filename) {
		var goVersion = getGolangVersion(filename)
		fmt.Printf("当前下载版本: %s\n", goVersion)
		utils.DownloadFile(downloadURL, "./"+filename)
		fmt.Println(`下载完成`)
	}

	go utils.Open(filename)

}

// 获取下载连接(golang)
//
// 返回 (下载地址, 文件名)
func getGolangDownloadURL() (string, string) {
	var gq, err = goquery.NewDocument(goMirrorURL + "/dl")
	if err != nil {
		panic(err)
	}
	var lists = gq.Find(".download.downloadBox")
	var now string = ""
	lists.Each(func(i int, s *goquery.Selection) {
		if now != "" {
			return
		}
		var text = s.Text()
		if strings.Contains(text, "windows") {
			now, _ = s.Attr("href")
		}
	})
	var output = goMirrorURL + now
	return output, filepath.Base(now)
}

// 获取`getGolangDownloadURL`版本号
//
// 其实就是格式化一下啦
//
// go1.15.6.windows-amd64.msi => go1.15.6
func getGolangVersion(rawstring string) string {
	var goVersion = filepath.Base(rawstring) // go1.15.6.windows-amd64.msi
	goVersion = strings.Split(goVersion, `.windows-amd64.msi`)[0]
	return goVersion
}

// 诶, 不知道安装哪个版本
func installGIT() {
	var filename = filepath.Base(gitDownloadURL)
	if !utils.Check(filename) {
		fmt.Printf("当前下载: %s\n", filename)
		utils.DownloadFile(gitDownloadURL, filename)
	}
	go utils.Open(filename)
}

func installVSCODE() {
	go utils.Open(vscodeBatFile)
}

func init() {
	var constTaobaoMirrorURL = `http://npm.taobao.org/mirrors/git-for-windows/`
	if utils.Arch() == utils.Bit64 {
		gitDownloadURL = constTaobaoMirrorURL + `v2.30.0-rc0.windows.1/Git-2.30.0-rc0-64-bit.exe`
	} else {
		gitDownloadURL = constTaobaoMirrorURL + `v2.30.0-rc0.windows.1/Git-2.30.0-rc0-32-bit.exe`
	}
	err := RestoreAsset(".", vscodeBatFile)
	if err != nil {
		panic(err)
	}
}

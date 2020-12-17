package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"

	"github.com/jsxxzy/ngo/utils"
)

func main() {
	fmt.Println("are you ok?")
}

func setup() {
	install_go()
	install_vscode()
	install_git()
}

func install_go() {

	// ######## 创建环境变量
	utils.SetGlobalENVBool("GO111MODULE", "on")
	utils.SetGlobalENVBool("GOPROXY", "https://goproxy.io,direct")
	// ########

}

func getGolangDownloadURL() string {
	var gq, err = goquery.NewDocument("https://studygolang.com/dl")
	if err != nil {
		panic(err)
	}
	ctxs := gq.Find(".download.downloadBox")
	return ""
}

func install_git() {

}

func install_vscode() {

}

package utils

import (
	"errors"
	"os/exec"
	"strings"
)

// SetGlobalENVBool 设置全局环境变量
func SetGlobalENVBool(envName, data string) bool {
	return SetGlobalENV(envName, data) == nil
}

// SetGlobalENV 设置全局环境变量
//
// 执行该命令可能并不能正确的判断是否成功
//
// `windows` 的操作系统相关东西不太懂...
func SetGlobalENV(envName, data string) error {
	var cmd = exec.Command("SETX", envName, data)
	err := cmd.Run()
	if err != nil {
		return err
	}
	outData, err := GetGlobalENV(envName)
	outData = strings.TrimSpace(outData)
	if err != nil {
		return err
	}
	var midFlag = outData == data
	if midFlag {
		return nil
	}
	return errors.New("设置失败, 比对的结果不相同")
}

// GetGlobalENV 获取某个全局环境变量
func GetGlobalENV(envName string) (string, error) {
	var input = "%" + envName + "%"
	var cmd = exec.Command("ECHO", input)
	dataByte, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	var data = string(dataByte)
	return data, nil
}

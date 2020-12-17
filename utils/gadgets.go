package utils

import (
	"os"

	"github.com/jsxxzy/ngo/utils/open"
)

// BitType 操作系统位数类型
type BitType int

const (
	// Bit32 32位操作系统
	Bit32 BitType = 0
	// Bit64 64位操作系统
	Bit64 BitType = 1
)

// Arch 获取操作系统位数
//
// 不造写的对不对, 将就将就
//
// 还要啥自行车啊
func Arch() BitType {
	bit := 32 << (^uint(0) >> 63)
	switch bit {
	case 32:
		return Bit32
	case 64:
		return Bit64
	default:
		return Bit32
	}
}

// Open 打开文件
//
// 主要是为了打开 `.exe` | `.msi` 可执行文件
func Open(runBinFilePath string) error {
	return open.Start(runBinFilePath)
}

// Check 判断文件是否存在
func Check(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

package variable

import (
	"log"
	"os"
)

var (
	BasePath string // 定义项目的根目录
)

func init() {
	// 1.初始化程序根目录
	if path, err := os.Getwd(); err == nil {
		BasePath = path
	} else {
		log.Fatal("读取根目录时出错！")
	}
}

package utils

import (
	"log"
)

func init() {

	log.SetFlags(0)
}

// 日志 打印
func LogPrint(format string, values ...interface{}) {

	log.Printf("[RPC-debug] "+format, values...)
}

// 错误 打印
func LogPrintError(err error) {

	if err != nil {
		LogPrint("[ERROR] %v\n", err)
	}
}

// 错误 打印&&退出
func LogFatalfError(err error) {

	if err != nil {
		log.Fatalf("[ERROR] %v\n", err)
	}
}

// 错误 转异常抛出
func ErrToPanic(e error) {
	if e != nil {
		panic(e)
	}
}



package bootstrap

import (
	"github.com/donnie4w/go-logger/logger"
)

func InitLog() {
	//设置日志打印文件夹和文件
	logger.SetRollingDaily("logs", "gome_logs.log")
	//设置为控制台输出
	logger.SetConsole(true)
	//设置日志级别为debug
	logger.SetLevel(logger.DEBUG)
	//打印日志
	logger.Debug("日志设置成功！")
}

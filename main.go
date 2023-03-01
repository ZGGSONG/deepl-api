package main

import (
	"deepl_api/server"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// 初始化日志
	logger := &lumberjack.Logger{
		Filename:   "./Log/deepl_api.log",
		MaxSize:    10,   // 日志文件大小，单位是 MB
		MaxBackups: 3,    // 最大过期日志保留个数
		MaxAge:     28,   // 保留过期文件最大时间，单位 天
		Compress:   true, // 是否压缩日志，默认是不压缩。这里设置为true，压缩日志
		LocalTime:  true, // 是否使用本地时间，默认是使用UTC时间
	}
	log.SetOutput(logger) // logrus 设置日志的输出方式
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	server.Run()
}

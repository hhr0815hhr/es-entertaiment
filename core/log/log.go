package log

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log, elog *zap.SugaredLogger

func InitLogger() {
	writeSyncer := getLogWriter()
	eWriteSyncer := getELogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	eCore := zapcore.NewCore(encoder, eWriteSyncer, zapcore.DebugLevel)

	log = zap.New(core, zap.AddCaller()).Sugar()
	elog = zap.New(eCore, zap.AddCaller()).Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	path, _ := os.Getwd()
	lumberjackLogger := &lumberjack.Logger{
		Filename:   path + "/runtime/log/info.log", // 日志文件路径
		MaxSize:    10,                             // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 5,                              // 日志文件按日期切割时，最多保留的文件个数
		MaxAge:     30,                             // 文件最多保存多少天
		Compress:   true,                           // 是否压缩旧文件
	}
	return zapcore.AddSync(lumberjackLogger)
}

func getELogWriter() zapcore.WriteSyncer {
	path, _ := os.Getwd()
	lumberjackLogger := &lumberjack.Logger{
		Filename:   path + "/runtime/log/error.log", // 日志文件路径
		MaxSize:    10,                              // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 5,                               // 日志文件按日期切割时，最多保留的文件个数
		MaxAge:     30,                              // 文件最多保存多少天
		Compress:   true,                            // 是否压缩旧文件
	}
	return zapcore.AddSync(lumberjackLogger)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(template string, args ...interface{}) {
	log.Infof(template, args...)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	log.Debugf(template, args...)
}

func Error(args ...interface{}) {
	elog.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	elog.Errorf(template, args...)
}

func Fatal(args ...interface{}) {
	elog.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	elog.Fatalf(template, args...)
}

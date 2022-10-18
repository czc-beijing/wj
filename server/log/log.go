package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func NewLog() {
	hook := lumberjack.Logger{
		Filename: "./logs/service.log",
		MaxSize:  1024,
		MaxAge:   8,
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		MessageKey:     "message",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)
	multiWrite := zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook))
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		multiWrite,
		atomicLevel,
	)
	Logger = zap.New(core)
}

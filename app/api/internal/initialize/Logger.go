package initialize

import (
	"gocument/app/api/global"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func SetupLogger() {
	//
	level := zap.NewAtomicLevel()
	level.SetLevel(zapcore.DebugLevel)

	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:       "message",
		LevelKey:         "level",
		TimeKey:          "time",
		NameKey:          "logger",
		CallerKey:        "caller",
		StacktraceKey:    "stacktrace",
		LineEnding:       zapcore.DefaultLineEnding, //行结束符
		EncodeLevel:      zapcore.CapitalColorLevelEncoder,
		EncodeTime:       zapcore.ISO8601TimeEncoder,
		EncodeDuration:   zapcore.StringDurationEncoder,
		EncodeCaller:     zapcore.FullCallerEncoder,
		ConsoleSeparator: " ",
	})
	cores := [...]zapcore.Core{
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level),
		zapcore.NewCore(encoder, zapcore.AddSync(getwritesync()), level),
	}
	global.Logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())
	defer func(Logger *zap.Logger) {
		_ = Logger.Sync()
	}(global.Logger)
	global.Logger.Info("initialize logger success")
}
func getwritesync() zapcore.WriteSyncer {
	lumberJACKLogger := &lumberjack.Logger{
		Filename:   global.Config.ZapConfig.Filename,
		MaxSize:    global.Config.ZapConfig.MaxSize,
		MaxAge:     global.Config.ZapConfig.MaxAge,
		MaxBackups: global.Config.ZapConfig.MaxBackups,
		LocalTime:  true,
	}
	return zapcore.AddSync(lumberJACKLogger)
}

package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	chainCfg "github.com/KuChainNetwork/kuchain/chain/config"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/spf13/viper"
	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/cli"
	tmflags "github.com/tendermint/tendermint/libs/cli/flags"
)

// SetLoggerToContext set logger to ctx
func SetLoggerToContext(ctx *server.Context) error {
	zapLogger := mkZapLogger(viper.GetBool(cli.TraceFlag))

	// process log level for cosmos-sdk
	logLvCfg := viper.GetString("log_level")
	logger, err := tmflags.ParseLogLevel(logLvCfg, NewLogger(zapLogger), cfg.DefaultLogLevel())
	if err != nil {
		return err
	}

	ctx.Logger = logger.With("module", "main")
	ctx.Config = chainCfg.DefaultConfig()
	return nil
}

func mkZapLogger(isDebug bool) *zap.Logger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "tm",
		LevelKey:       "lv",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     zapcore.EpochNanosTimeEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	if isDebug {
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	}

	config := zap.NewDevelopmentConfig()

	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel) // most small
	config.EncoderConfig = encoderConfig
	config.Development = isDebug

	logger, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("zap logger build err by %s", err.Error()))
	}

	return logger.WithOptions(zap.AddCallerSkip(2))
}

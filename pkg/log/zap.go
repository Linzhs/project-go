package log

import (
	"go.uber.org/zap"
)

//import "gopkg.in/natefinch/lumberjack.v2"

func NewZapLogger() {
	//cfg := zap.Config{
	//	Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
	//	Encoding:         "json",
	//	OutputPaths:      []string{"stdout"},
	//	ErrorOutputPaths: []string{"stderr"},
	//	InitialFields:    map[string]interface{}{},
	//	EncoderConfig:    nil,
	//}
	//logger, err := cfg.Build()
	//if err != nil {
	//	panic(err)
	//}

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.Sugar()
	zap.S()
}
